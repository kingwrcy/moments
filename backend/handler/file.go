package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type FileHandler struct {
	base BaseHandler
}

func NewFileHandler(injector do.Injector) *FileHandler {
	return &FileHandler{do.MustInvoke[BaseHandler](injector)}
}

func (f FileHandler) Get(c echo.Context) error {
	filename := c.Param("filename")
	f.base.log.Info().Msgf("Getting file %s", filename)
	if filename == "" {
		return c.HTML(404, "not found")
	}
	fp := path.Join(f.base.cfg.UploadDir, filename)
	if _, err := os.Stat(fp); errors.Is(err, os.ErrNotExist) {
		return c.HTML(404, "not found")
	}
	return c.File(path.Join(f.base.cfg.UploadDir, filename))
}

// Upload godoc
//
//	@Tags		File
//	@Summary	上传图片
//	@Accept		json
//	@Produce	json
//	@Param		x-api-token	header	string	true	"登录TOKEN"
//	@Success	200
//	@Router		/api/file/upload [post]
func (f FileHandler) Upload(c echo.Context) error {
	var (
		result []string
	)
	form, err := c.MultipartForm()
	if err != nil {
		f.base.log.Error().Msgf("读取上传图片异常:%s", err)
		return FailRespWithMsg(c, Fail, "上传图片异常")
	}
	files := form.File["files"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			f.base.log.Error().Msgf("打开上传图片异常:%s", err)
			return FailRespWithMsg(c, Fail, "上传图片异常")
		}
		// Destination
		filename := strings.ReplaceAll(uuid.NewString(), "-", "")
		path := path.Join(f.base.cfg.UploadDir, filename)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			f.base.log.Error().Msgf("创建父级目录异常:%s", err)
			return FailRespWithMsg(c, Fail, "创建父级目录异常")
		}
		dst, err := os.Create(path)
		if err != nil {
			f.base.log.Error().Msgf("打开目标图片异常:%s", err)
			return FailRespWithMsg(c, Fail, "上传图片异常")
		}
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			f.base.log.Error().Msgf("复制图片异常:%s", err)
			return FailRespWithMsg(c, Fail, "上传图片异常")
		}

		src.Close()
		dst.Close()
		result = append(result, "/upload/"+filename)
	}
	return SuccessResp(c, result)
}

type PreSignedReq struct {
	ContentType string `json:"contentType,omitempty"` //图片mime类型
}

type s3PresignedResp struct {
	PreSignedUrl string `json:"preSignedUrl,omitempty"` //S3预签名上传地址
	ImageUrl     string `json:"imageUrl,omitempty"`     //实际的图片地址
}

// S3PreSigned godoc
//
//	@Tags		File
//	@Summary	S3预签名
//	@Accept		json
//	@Produce	json
//	@Param		object		body		PreSignedReq	true	"S3预签名"
//	@Param		x-api-token	header		string			true	"登录TOKEN"
//	@Success	200			{object}	s3PresignedResp
//	@Router		/api/file/s3PreSigned [post]
func (f FileHandler) S3PreSigned(c echo.Context) error {

	var (
		req         PreSignedReq
		sysConfig   db.SysConfig
		sysConfigVo vo.FullSysConfigVO
	)
	if err := c.Bind(&req); err != nil {
		return FailResp(c, ParamError)
	}

	if err := f.base.db.First(&sysConfig).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailResp(c, Fail)
	}
	if err := json.Unmarshal([]byte(sysConfig.Content), &sysConfigVo); err != nil {
		f.base.log.Error().Msgf("无法反序列化系统配置, %s", err)
		return FailRespWithMsg(c, Fail, err.Error())
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(sysConfigVo.S3.Region),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{URL: sysConfigVo.S3.Endpoint}, nil
		})),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(sysConfigVo.S3.AccessKey, sysConfigVo.S3.SecretKey, "")))
	if err != nil {
		f.base.log.Error().Msgf("无法加载SDK配置, %s", err)
		return FailRespWithMsg(c, Fail, err.Error())
	}

	client := s3.NewFromConfig(cfg)
	presignedClient := s3.NewPresignClient(client)

	key := fmt.Sprintf("moments/%s/%s", time.Now().Format("2006/01/02"), strings.ReplaceAll(uuid.NewString(), "-", ""))
	presignedResult, err := presignedClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(sysConfigVo.S3.Bucket),
		Key:         aws.String(key),
		ContentType: aws.String(req.ContentType),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Minute * 5
	})

	if err != nil {
		f.base.log.Error().Msgf("无法获取预签名URL, %s", err)
		return FailRespWithMsg(c, Fail, fmt.Sprintf("无法获取预签名URL, %s", err))
	}
	suffix := sysConfigVo.S3.ThumbnailSuffix
	if !strings.HasPrefix(suffix, "?") {
		suffix = "?" + suffix
	}
	return SuccessResp(c, s3PresignedResp{
		PreSignedUrl: presignedResult.URL,
		ImageUrl:     fmt.Sprintf("%s/%s%s", sysConfigVo.S3.Domain, key, suffix),
	})
}
