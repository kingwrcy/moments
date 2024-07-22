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
	f.base.log.Info().Msgf("filename :%s", filename)
	if filename == "" {
		return c.HTML(404, "not found")
	}
	filepath := path.Join(f.base.cfg.UploadDir, filename)
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return c.HTML(404, "not found")
	}
	return c.File(path.Join(f.base.cfg.UploadDir, filename))
}

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
		filename := strings.ReplaceAll(uuid.New().String(), "-", "")
		filepath := path.Join(f.base.cfg.UploadDir, filename)
		dst, err := os.Create(filepath)
		if err != nil {
			f.base.log.Error().Msgf("打开目标图片异常:%s", err)
			return FailRespWithMsg(c, Fail, "上传图片异常")
		}
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			f.base.log.Error().Msgf("负责图片异常:%s", err)
			return FailRespWithMsg(c, Fail, "上传图片异常")
		}

		src.Close()
		dst.Close()
		result = append(result, "/api/file/get/"+filename)
	}
	return SuccessResp(c, result)
}

type PreSignedReq struct {
	ContentType string `json:"contentType,omitempty"`
}

func (f FileHandler) S3PreSigned(c echo.Context) error {

	var (
		req         PreSignedReq
		sysConfig   db.SysConfig
		sysConfigVo vo.SysConfigVO
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
	return SuccessResp(c, h{
		"preSignedUrl": presignedResult.URL,
		"imageUrl":     fmt.Sprintf("%s/%s%s", sysConfigVo.S3.Domain, key, suffix),
	})
}
