package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/kingwrcy/moments/db"
	fs_util "github.com/kingwrcy/moments/util"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"golang.org/x/net/html"
	"gorm.io/gorm"
)

type MemoHandler struct {
	base BaseHandler
	hc   http.Client
}

func NewMemoHandler(injector do.Injector) *MemoHandler {
	return &MemoHandler{
		base: do.MustInvoke[BaseHandler](injector),
		hc:   http.Client{},
	}
}

type memoListResp struct {
	List    []db.Memo `json:"list,omitempty"`    //memo列表
	HasNext bool      `json:"hasNext,omitempty"` //是否有下一页
	Total   int64     `json:"total,omitempty"`   //总数
}

/*
这里通过 memo.Imgs 生成 ImgConfigs，ImgConfig 包含图片的原始 Url 和 ThumbUrl
正常情况下这里应该是在上传图片的时候来保存 ImgConfig，而不是在获取 memo 时计算
但是由于历史实现的 memo 模型结构不支持保存额外的图片信息，所以这里先临时实现一个版本，有精力再重构
*/
func (m MemoHandler) handleImgConfigs(sysConfigVO *vo.FullSysConfigVO, memo *db.Memo) {
	var imgConfigs []*vo.ImgConfig

	imgs := strings.Split(memo.Imgs, ",")
	for _, img := range imgs {
		if img == "" {
			continue
		}

		imgConfig := &vo.ImgConfig{
			Url:      &img,
			ThumbUrl: &img,
		}

		if strings.HasPrefix(img, "/upload/") {
			thumb_filename := img + "_thumb"
			thumb_filepath := path.Join(m.base.cfg.UploadDir, path.Base(thumb_filename))
			if fs_util.Exists(thumb_filepath) {
				imgConfig.ThumbUrl = &thumb_filename
			}
		} else if sysConfigVO.S3.ThumbnailSuffix != "" && strings.HasPrefix(img, sysConfigVO.S3.Domain) {
			thumbnailSuffix := sysConfigVO.S3.ThumbnailSuffix
			if !strings.HasPrefix(thumbnailSuffix, "?") {
				thumbnailSuffix = "?" + thumbnailSuffix
			}

			newThumbUrl := img + thumbnailSuffix
			imgConfig.ThumbUrl = &newThumbUrl
		}

		imgConfigs = append(imgConfigs, imgConfig)
	}

	memo.ImgConfigs = &imgConfigs
}

// ListMemos godoc
//
//	@Tags		Memo
//	@Summary	查询memos列表
//	@Accept		json
//	@Produce	json
//	@Param		object	body		vo.ListMemoReq	true	"查询memos列表"
//	@Success	200		{object}	memoListResp
//	@Router		/api/memo/list [post]
func (m MemoHandler) ListMemos(c echo.Context) error {
	var (
		req         vo.ListMemoReq
		list        []db.Memo
		total       int64
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
	)

	ctx := c.(CustomContext)
	currentUser := ctx.CurrentUser()
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	m.base.db.First(&sysConfig)
	_ = json.Unmarshal([]byte(sysConfig.Content), &sysConfigVO)
	offset := (req.Page - 1) * req.Size

	tx := m.base.db.Model(&db.Memo{}).Preload("User", func(x *gorm.DB) *gorm.DB {
		return x.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl")
	})

	if req.Start != nil {
		tx = tx.Where("createdAt >= ?", req.Start)
	}
	if req.End != nil {
		tx = tx.Where("createdAt <= ?", req.End)
	}
	if req.ContentContains != "" {
		tx = tx.Where("content like ?", "%"+req.ContentContains+"%")
	}
	if req.ShowType != nil && *req.ShowType >= 0 {
		tx = tx.Where("showType=?", req.ShowType)
	}
	if currentUser == nil {
		tx = tx.Where("showType = 1")
		tx = tx.Where("createdAt <= ?", time.Now())
	} else {
		tx = tx.Where("userId = ? or (userId <> ? and showType = 1)", currentUser.Id, currentUser.Id)
		tx = tx.Where("userId = ? or createdAt <= ?", currentUser.Id, time.Now())
	}
	if req.Tag != "" {
		if strings.Contains(req.Tag, ",") {
			tags := strings.Split(req.Tag, ",")
			for _, tag := range tags {
				tx = tx.Where("tags like ?", fmt.Sprintf("%%%s,%%", tag))
			}
		} else {
			tx = tx.Where("tags like ?", fmt.Sprintf("%%%s,%%", req.Tag))
		}
	}
	if req.Username != "" {
		var target db.User
		if err := m.base.db.Where("username = ?", req.Username).First(&target).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return FailRespWithMsg(c, Fail, "不存在的用户")
		}
		tx = tx.Where("userId = ?", target.Id)
	}
	if req.UserId != nil {
		tx = tx.Where("userId = ?", req.UserId)
	}
	tx.Session(&gorm.Session{}).Order("pinned desc, createdAt desc").Limit(req.Size).Offset(offset).Find(&list)
	tx.Session(&gorm.Session{}).Count(&total)

	for i, memo := range list {
		var comments []db.Comment
		m.base.db.Where("memoId = ?", memo.Id).Order(fmt.Sprintf("createdAt %s", sysConfigVO.CommentOrder)).Limit(5).Find(&comments)
		list[i].Comments = comments
	}

	for i := range list {
		m.handleImgConfigs(&sysConfigVO, &list[i])
	}

	return SuccessResp(c, memoListResp{
		List:    list,
		Total:   total,
		HasNext: int64(req.Page*req.Size) < total,
	})
}

// RemoveMemo godoc
//
//	@Tags		Memo
//	@Summary	删除memos
//	@Accept		json
//	@Produce	json
//	@Param		id			query	int		true	"memoID"
//	@Param		x-api-token	header	string	true	"登录TOKEN"
//	@Success	200
//	@Router		/api/memo/remove [post]
func (m MemoHandler) RemoveMemo(c echo.Context) error {
	ctx := c.(CustomContext)
	currentUser := ctx.CurrentUser()
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}
	var (
		memo db.Memo
	)
	if err = m.base.db.First(&memo, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailResp(c, ParamError)
	}

	if currentUser.Id != memo.UserId && currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "没有权限")
	}
	if m.base.db.Delete(&memo).RowsAffected != 1 {
		return FailRespWithMsg(c, Fail, "删除失败")
	}

	return SuccessResp(c, h{})
}

// LikeMemo godoc
//
//	@Tags		Memo
//	@Summary	点赞memo
//	@Accept		json
//	@Produce	json
//	@Param		id	query	int	true	"memoID"
//	@Success	200
//	@Router		/api/memo/like [post]
func (m MemoHandler) LikeMemo(c echo.Context) error {
	var (
		memo        db.Memo
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
		token       string
	)
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	m.base.db.First(&sysConfig)
	_ = json.Unmarshal([]byte(sysConfig.Content), &sysConfigVO)

	if sysConfigVO.EnableGoogleRecaptcha {
		token = c.QueryParam("token")
		if token == "" {
			return FailRespWithMsg(c, ParamError, "token不能为空")
		}
		if err := checkGoogleRecaptcha(m.base.log, sysConfigVO, token); err != nil {
			return FailRespWithMsg(c, Fail, err.Error())
		}
	}
	if err = m.base.db.First(&memo, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailResp(c, ParamError)
	}
	memo.FavCount = memo.FavCount + 1
	if m.base.db.Updates(&memo).RowsAffected != 1 {
		return FailRespWithMsg(c, Fail, "点赞失败")
	}
	return SuccessResp(c, h{})
}

// FindAndReplaceTags 处理 markdown 文本
func FindAndReplaceTags(content string) (string, []string) {
	// 定义正则表达式来匹配标签
	tagRegex := regexp.MustCompile(`#[^\s#,]+`)
	headingRegex := regexp.MustCompile(`^#{1,6}\s+.+`)

	// 使用正则表达式找到所有标签
	tags := tagRegex.FindAllString(content, -1)

	// 去掉重复的标签，并移除 # 号
	tagMap := make(map[string]bool)
	var uniqueTags []string
	for _, tag := range tags {
		trimmedTag := strings.Trim(tag, "#")
		if trimmedTag != "" && !tagMap[trimmedTag] {
			tagMap[trimmedTag] = true
			uniqueTags = append(uniqueTags, trimmedTag)
		}
	}

	// 替换标签为空，并清理多余的逗号和空格，但不影响标题
	replacedContent := strings.Split(content, "\n")
	for i, line := range replacedContent {
		if !headingRegex.MatchString(line) {
			// 替换标签为空
			line = tagRegex.ReplaceAllString(line, "")
			// 清理多余的逗号和空格
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, ",", "")
			replacedContent[i] = line
		}
	}

	return strings.Join(replacedContent, "\n"), uniqueTags
}

// SaveMemo godoc
//
//	@Tags		Memo
//	@Summary	保存memos
//	@Accept		json
//	@Produce	json
//	@Param		x-api-token	header	string			true	"登录TOKEN"
//	@Param		object		body	vo.SaveMemoReq	true	"保存memos"
//	@Success	200
//	@Router		/api/memo/save [post]
func (m MemoHandler) SaveMemo(c echo.Context) error {
	var (
		req     vo.SaveMemoReq
		extJson string
	)
	err := c.Bind(&req)
	if err != nil {
		m.base.log.Error().Msgf("保存memo时参数校验失败,原因:%s", err)
		return FailResp(c, ParamError)
	}

	var memo db.Memo
	var now = time.Now()
	ctx := c.(CustomContext)
	currentUser := ctx.CurrentUser()

	if req.ID > 0 {
		if err = m.base.db.First(&memo, req.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return FailResp(c, ParamError)
		}
		if memo.UserId != currentUser.Id {
			return FailResp(c, ParamError)
		}
		memo.UpdatedAt = &now
	} else {
		memo.CreatedAt = &now
		memo.UserId = currentUser.Id
		memo.FavCount = 0
		memo.CommentCount = 0
	}

	//content, tags := FindAndReplaceTags(req.Content)
	//m.base.log.Info().Msgf("tags is %+v,content is %s", tags, content)
	if len(req.Tags) == 0 {
		memo.Tags = nil
	} else {
		memoTags := strings.Join(req.Tags, ",")
		if memoTags != "" {
			memoTags = fmt.Sprintf("%s,", memoTags)
			memo.Tags = &memoTags
		}
	}
	memo.Content = strings.TrimSpace(req.Content)

	bytes, _ := json.Marshal(req.Ext)
	extJson = string(bytes)

	memo.Imgs = strings.Join(req.Imgs, ",")
	memo.Location = req.Location
	memo.ExternalUrl = req.ExternalUrl
	memo.ExternalTitle = req.ExternalTitle
	memo.ExternalFavicon = req.ExternalFavicon
	memo.Pinned = req.Pinned
	memo.Ext = extJson
	memo.ShowType = req.ShowType

	// 处理自定义时间
	if req.CreatedAt != nil {
		*memo.CreatedAt = req.CreatedAt.Local()
	}

	m.base.db.Save(&memo)

	return SuccessResp(c, h{})
}

// GetMemo godoc
//
//	@Tags		Memo
//	@Summary	获取单个memo详情
//	@Accept		json
//	@Produce	json
//	@Param		id	query		int	true	"memoID"
//	@Success	200	{object}	db.Memo
//	@Router		/api/memo/get [post]
func (m MemoHandler) GetMemo(c echo.Context) error {
	var (
		memo        db.Memo
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
	)

	ctx := c.(CustomContext)
	currentUser := ctx.CurrentUser()

	m.base.db.First(&sysConfig)

	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	latest := c.QueryParam("latest")

	if err = m.base.db.Preload("User", func(x *gorm.DB) *gorm.DB {
		return x.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl")
	}).First(&memo, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailResp(c, ParamError)
	}

	if *memo.ShowType != 1 && (currentUser == nil || currentUser.Id != memo.UserId) {
		return FailRespWithMsg(c, Fail, "暂无权限查看")
	}

	var comments []db.Comment
	tx := m.base.db.Where("memoId = ?", memo.Id).Order("createdAt DESC")
	if latest != "" {
		tx.Limit(5)
	}
	tx.Find(&comments)

	memo.Comments = comments

	m.handleImgConfigs(&sysConfigVO, &memo)

	return SuccessResp(c, memo)
}

// SetPinned godoc
//
//	@Tags		Memo
//	@Summary	置顶单个memo
//	@Accept		json
//	@Produce	json
//	@Param		id			query	int		true	"memoID"
//	@Param		x-api-token	header	string	true	"登录TOKEN"
//	@Success	200
//	@Router		/api/memo/setPinned [post]
func (m MemoHandler) SetPinned(c echo.Context) error {
	ctx := c.(CustomContext)
	currentUser := ctx.CurrentUser()
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}
	var (
		memo db.Memo
	)
	if err = m.base.db.First(&memo, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailResp(c, ParamError)
	}

	if currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "没有权限")
	}

	m.base.db.Table("Memo").Where("pinned = true").Update("pinned", false)
	pinned := *memo.Pinned
	if err = m.base.db.Table("Memo").Where("id=?", id).Update("pinned", !pinned).Error; err != nil {
		return FailRespWithMsg(c, Fail, err.Error())
	}
	return SuccessResp(c, h{})
}

type externalWebsite struct {
	Favicon string `json:"favicon,omitempty"` //favicon
	Title   string `json:"title,omitempty"`   //标题
}

// GetFaviconAndTitle godoc
//
//	@Tags		Memo
//	@Summary	获取外部链接的相关信息
//	@Accept		json
//	@Produce	json
//	@Param		url			query		string	true	"memoID"
//	@Param		x-api-token	header		string	true	"登录TOKEN"
//	@Success	200			{object}	externalWebsite
//	@Router		/api/memo/getFaviconAndTitle [post]
func (m MemoHandler) GetFaviconAndTitle(c echo.Context) error {
	websiteURL := c.QueryParam("url")
	favicon, title, err := getFaviconAndTitle(websiteURL)
	if err != nil {
		return FailResp(c, ParamError)
	}
	return SuccessResp(c, externalWebsite{
		Favicon: favicon,
		Title:   title,
	})
}

// getFaviconAndTitle tries to get the favicon URL and title from the given website URL.
func getFaviconAndTitle(websiteURL string) (string, string, error) {
	// Parse the provided URL to extract the domain
	parsedURL, err := url.Parse(websiteURL)
	if err != nil {
		return "", "", err
	}

	// Construct the base URL
	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	// Parse the HTML to find <link rel="icon"> or <link rel="shortcut icon"> and the <title>
	resp, err := http.Get(websiteURL)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", "", err
	}

	var favicon, title string
	var parseHTML func(*html.Node)
	parseHTML = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "link" {
				isIcon := false
				href := ""
				for _, attr := range n.Attr {
					if attr.Key == "rel" && (attr.Val == "icon" || attr.Val == "shortcut icon") {
						isIcon = true
					}
					if attr.Key == "href" {
						href = attr.Val
					}
				}
				if isIcon && href != "" {
					favicon = href
				}
			} else if n.Data == "title" && n.FirstChild != nil {
				title = n.FirstChild.Data
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseHTML(c)
		}
	}
	parseHTML(doc)

	// If favicon is found, resolve it to an absolute URL
	if favicon != "" {
		faviconURL, err := url.Parse(favicon)
		if err != nil {
			return "", "", err
		}
		faviconURL = parsedURL.ResolveReference(faviconURL)
		favicon = faviconURL.String()
	} else {
		// Try to get the favicon from /favicon.ico
		faviconURL := baseURL + "/favicon.ico"
		resp, err := http.Get(faviconURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			favicon = faviconURL
			resp.Body.Close()
		} else if resp != nil {
			resp.Body.Close()
		}
	}

	return favicon, title, nil
}

// GetDoubanMovieInfo godoc
//
//	@Tags       Memo
//	@Summary    获取豆瓣电影详情
//	@Accept     json
//	@Produce    json
//	@Param      id          query       int     true    "豆瓣电影ID"
//	@Param      x-api-token header      string  true    "登录TOKEN"
//	@Success    200         {object}    vo.DoubanMovie
//	@Router     /api/memo/getDoubanMovieInfo [post]
func (m MemoHandler) GetDoubanMovieInfo(c echo.Context) error {

	var (
		book        vo.DoubanMovie
		sysConfigVo vo.FullSysConfigVO
		sysConfig   db.SysConfig
		userAgent   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	)

	if err := m.base.db.First(&sysConfig).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailRespWithMsg(c, Fail, "系统配置为空")
	}
	err := json.Unmarshal([]byte(sysConfig.Content), &sysConfigVo)
	if err != nil {
		return FailRespWithMsg(c, Fail, "读取系统配置异常")
	}

	id := c.QueryParam("id")
	target := fmt.Sprintf("https://movie.douban.com/subject/%s/", id)

	req, _ := http.NewRequest("GET", target, nil)
	req.Header.Set("User-Agent", userAgent)
	start := time.Now()
	res, err := m.hc.Do(req)
	m.base.log.Info().Str("豆瓣读书ID", id).Str("URL", target).Str("耗时", fmt.Sprintf("%f秒", time.Since(start).Seconds())).Msgf("获取豆瓣读书")
	if err != nil {
		m.base.log.Error().Msgf("获取豆瓣电影异常:%s", err.Error())
		return FailRespWithMsg(c, Fail, err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		m.base.log.Error().Msgf("豆瓣电影API返回码不是200,而是:%d", res.StatusCode)
		return FailRespWithMsg(c, Fail, fmt.Sprintf("豆瓣读书API返回码不是200,而是:%d,URL:%s", res.StatusCode, target))
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		m.base.log.Error().Msgf("初始化html错误,%s", err.Error())
		return FailRespWithMsg(c, Fail, fmt.Sprintf("初始化html错误,%s", err.Error()))
	}

	doc.Find("meta[property]").Each(func(i int, s *goquery.Selection) {
		if properties, exists := s.Attr("property"); exists {
			value := s.AttrOr("content", "")
			if properties == "og:title" {
				book.Title = value
			} else if properties == "og:description" {
				book.Desc = value
			} else if properties == "og:image" {
				book.Image = value
			} else if properties == "video:director" {
				book.Director = value
			} else if properties == "video:actor" {
				book.Actors = value + "/"
			}
		}
	})
	book.Url = target
	book.ReleaseDate = doc.Find("span[property='v:initialReleaseDate']").AttrOr("content", "")
	book.Runtime = doc.Find("span[property='v:runtime']").AttrOr("content", "")
	book.Rating = doc.Find("strong.rating_num").Text()
	if book.Rating == "" {
		book.Rating = "未知评分"
	}

	if !strings.HasPrefix(book.Image, "http") {
		return FailRespWithMsg(c, Fail, "无法获取电影封面")
	}
	if sysConfigVo.EnableS3 {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(sysConfigVo.S3.Region),
			config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: sysConfigVo.S3.Endpoint}, nil
			})),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(sysConfigVo.S3.AccessKey, sysConfigVo.S3.SecretKey, "")))
		if err != nil {
			m.base.log.Error().Msgf("无法加载S3 SDK配置, %s", err)
			return FailRespWithMsg(c, Fail, err.Error())
		}
		imageReq, _ := http.NewRequest("GET", book.Image, nil)
		imageReq.Header.Set("User-Agent", userAgent)
		imageResponse, err := m.hc.Do(imageReq)
		if err != nil {
			return FailRespWithMsg(c, Fail, fmt.Sprintf("下载豆瓣电影图片异常:%s", err.Error()))
		}
		defer imageResponse.Body.Close()
		client := s3.NewFromConfig(cfg)
		key := fmt.Sprintf("moments/%s/%s", time.Now().Format("2006/01/02"), strings.ReplaceAll(uuid.NewString(), "-", ""))
		_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(sysConfigVo.S3.Bucket),
			Key:    aws.String(key),
			Body:   imageResponse.Body,
		})
		if err != nil {
			return FailRespWithMsg(c, Fail, fmt.Sprintf("上传图片到s3异常:%s", err.Error()))
		}
		book.Image = fmt.Sprintf("%s/%s", sysConfigVo.S3.Domain, key)
	} else {
		image, err := downloadImage(book.Image, m.base.log, m.base.cfg)
		if err != nil {
			return FailRespWithMsg(c, Fail, fmt.Sprintf("下载豆瓣电影图片异常:%s", err.Error()))
		}
		book.Image = image
	}
	return SuccessResp(c, book)
}

func downloadImage(src string, log zerolog.Logger, conf *vo.AppConfig) (string, error) {
	start := time.Now()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", src, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	response, err := client.Do(req)
	log.Info().Msgf("下载图片完成:%s,耗时:%f", src, time.Since(start).Seconds())
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	key := strings.ReplaceAll(uuid.NewString(), "-", "")
	filepath := fmt.Sprintf("%s/%s.jpg", conf.UploadDir, key)
	dst, err := os.Create(filepath)
	log.Info().Msgf("保存图片到本地完成:%s,耗时:%f", src, time.Since(start).Seconds())
	if err != nil {
		log.Error().Msgf("打开目标图片异常:%s", err)
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, response.Body)
	log.Info().Msgf("保存图片到本地完成:%s,耗时:%f", src, time.Since(start).Seconds())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/upload/%s.jpg", key), err
}

// GetDoubanBookInfo godoc
//
//	@Tags       Memo
//	@Summary    获取豆瓣读书详情
//	@Accept     json
//	@Produce    json
//	@Param      id          query       int     true    "豆瓣读书ID"
//	@Param      x-api-token header      string  true    "登录TOKEN"
//	@Success    200         {object}    vo.DoubanBook
//	@Router     /api/memo/getDoubanBookInfo [post]
func (m MemoHandler) GetDoubanBookInfo(c echo.Context) error {

	var (
		book        vo.DoubanBook
		sysConfigVo vo.FullSysConfigVO
		sysConfig   db.SysConfig
		re          = regexp.MustCompile(`\d{4}-\d{1,2}(-\d{1,2})?`)
		userAgent   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	)
	if err := m.base.db.First(&sysConfig).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailRespWithMsg(c, Fail, "系统配置为空")
	}
	err := json.Unmarshal([]byte(sysConfig.Content), &sysConfigVo)
	if err != nil {
		return FailRespWithMsg(c, Fail, "读取系统配置异常")
	}

	id := c.QueryParam("id")
	m.base.log.Info().Msgf("开始分析豆瓣图书,id:%s", id)

	target := fmt.Sprintf("https://book.douban.com/subject/%s/", id)
	// Request the HTML page.
	client := &http.Client{}
	start := time.Now()
	req, _ := http.NewRequest("GET", target, nil)
	m.base.log.Info().Msgf("请求豆瓣读书地址:%s,耗时:%f秒", target, time.Since(start).Seconds())
	req.Header.Set("User-Agent", userAgent)
	res, err := client.Do(req)
	if err != nil {
		m.base.log.Error().Msgf("获取豆瓣读书异常:%s", err.Error())
		return FailRespWithMsg(c, Fail, err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		m.base.log.Error().Msgf("豆瓣读书API返回码不是200,而是:%d", res.StatusCode)
		return FailRespWithMsg(c, Fail, fmt.Sprintf("豆瓣读书API返回码不是200,而是:%d,URL:%s", res.StatusCode, target))
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		m.base.log.Error().Msgf("初始化html错误,%s", err.Error())
		return FailRespWithMsg(c, Fail, fmt.Sprintf("初始化html错误,%s", err.Error()))
	}

	doc.Find("meta[property]").Each(func(i int, s *goquery.Selection) {
		if properties, exists := s.Attr("property"); exists {
			value := s.AttrOr("content", "")
			if properties == "og:title" {
				book.Title = value
			} else if properties == "og:description" {
				book.Desc = value
			} else if properties == "og:image" {
				book.Image = value
			} else if properties == "book:author" {
				book.Author = value
			} else if properties == "book:isbn" {
				book.Isbn = value
			}
		}
	})
	book.Url = target
	book.Keywords = doc.Find("meta[name='keywords']").AttrOr("content", "")
	date := re.FindString(book.Keywords)
	if date != "" {
		book.PubDate = date
	}
	book.Rating = doc.Find("strong.rating_num").Text()
	if strings.TrimSpace(book.Rating) == "" {
		book.Rating = "暂无"
	}

	if !strings.HasPrefix(book.Image, "http") {
		return FailRespWithMsg(c, Fail, "无法获取图书封面")
	}
	if sysConfigVo.EnableS3 {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(sysConfigVo.S3.Region),
			config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: sysConfigVo.S3.Endpoint}, nil
			})),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(sysConfigVo.S3.AccessKey, sysConfigVo.S3.SecretKey, "")))
		if err != nil {
			m.base.log.Error().Msgf("无法加载S3 SDK配置, %s", err)
			return FailRespWithMsg(c, Fail, err.Error())
		}
		imageReq, _ := http.NewRequest("GET", book.Image, nil)
		imageReq.Header.Set("User-Agent", userAgent)
		imageResponse, err := client.Do(imageReq)
		if err != nil {
			return FailRespWithMsg(c, Fail, fmt.Sprintf("下载豆瓣图片异常:%s", err.Error()))
		}
		defer imageResponse.Body.Close()
		s3Client := s3.NewFromConfig(cfg)
		key := fmt.Sprintf("moments/%s/%s", time.Now().Format("2006/01/02"), strings.ReplaceAll(uuid.NewString(), "-", ""))
		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(sysConfigVo.S3.Bucket),
			Key:    aws.String(key),
			Body:   imageResponse.Body,
		})
		if err != nil {
			return FailRespWithMsg(c, Fail, fmt.Sprintf("上传图片到s3异常:%s", err.Error()))
		}
		book.Image = fmt.Sprintf("%s/%s", sysConfigVo.S3.Domain, key)
	} else {
		image, err := downloadImage(book.Image, m.base.log, m.base.cfg)
		if err != nil {
			return FailRespWithMsg(c, Fail, fmt.Sprintf("下载豆瓣图片异常:%s", err.Error()))
		}
		book.Image = image
	}
	return SuccessResp(c, book)
}
// backend/handler/memo.go

// ... 在 handler 中添加以下两个方法 ...

// GetSteamGameInfo 获取 Steam 游戏详情
// @Router /api/memo/getSteamGameInfo [post]
// GetSteamGameInfo 获取 Steam 游戏详情
// @Router /api/memo/getSteamGameInfo [post]
func (m MemoHandler) GetSteamGameInfo(c echo.Context) error {
	id := c.QueryParam("id")
	// 设置语言为中文
	target := fmt.Sprintf("https://store.steampowered.com/app/%s/?l=schinese", id)

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"

	req, _ := http.NewRequest("GET", target, nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", "birthtime=0; lastagecheckage=1-0-1900")

	res, err := m.hc.Do(req)
	if err != nil {
		return FailRespWithMsg(c, Fail, err.Error())
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return FailRespWithMsg(c, Fail, "解析HTML失败")
	}

	var game vo.SteamGame
	game.AppId = id
	game.Url = target
	
	// [修复] 使用 First() 防止获取到重复的标题 (例如 "九日九日")
	game.Name = strings.TrimSpace(doc.Find(".apphub_AppName").First().Text())
	
	game.Image = doc.Find("img.game_header_image_full").AttrOr("src", "")
	game.Description = strings.TrimSpace(doc.Find(".game_description_snippet").Text())
	game.ReleaseDate = strings.TrimSpace(doc.Find(".date").First().Text())
	game.Price = strings.TrimSpace(doc.Find(".game_purchase_price").First().Text())
	if game.Price == "" {
		game.Price = strings.TrimSpace(doc.Find(".discount_final_price").First().Text())
	}

	if game.Name == "" {
		return FailRespWithMsg(c, Fail, "无法获取游戏信息，可能ID无效")
	}

	return SuccessResp(c, game)
}

// GetTmdbInfo 获取 TMDB 详情
// @Router /api/memo/getTmdbInfo [post]
func (m MemoHandler) GetTmdbInfo(c echo.Context) error {
	id := c.QueryParam("id")
	typeStr := c.QueryParam("type") // "movie" or "tv"
	if typeStr == "" {
		typeStr = "movie"
	}
	target := fmt.Sprintf("https://www.themoviedb.org/%s/%s?language=zh-CN", typeStr, id)
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	req, _ := http.NewRequest("GET", target, nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	res, err := m.hc.Do(req)
	if err != nil {
		return FailRespWithMsg(c, Fail, err.Error())
	}
	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return FailRespWithMsg(c, Fail, "解析HTML失败")
	}
	
	var item vo.TmdbItem
	item.Id = id
	item.Type = typeStr
	item.Url = target
	
	// 1. Meta信息
	item.Title = doc.Find("meta[property='og:title']").AttrOr("content", "")
	item.Overview = strings.TrimSpace(doc.Find("meta[property='og:description']").AttrOr("content", ""))
	item.PosterPath = doc.Find("meta[property='og:image']").AttrOr("content", "")
	
	// 2. 标题清理
	if strings.Contains(item.Title, " (") {
		parts := strings.Split(item.Title, " (")
		item.Title = parts[0]
	}
	
	// 3. 日期
	dateStr := strings.TrimSpace(doc.Find(".facts .release").Text())
	if dateStr != "" {
		item.ReleaseDate = strings.TrimSpace(strings.Split(dateStr, " (")[0])
	} else {
		item.ReleaseDate = strings.TrimSpace(doc.Find(".tag.release_date").Text())
	}
	
	// 4. 评分
	scoreStr := doc.Find(".user_score_chart").AttrOr("data-percent", "")
	if scoreStr != "" {
		if scoreFloat, err := strconv.ParseFloat(scoreStr, 64); err == nil {
			item.VoteAverage = fmt.Sprintf("%.1f", scoreFloat/10.0)
		} else {
			item.VoteAverage = scoreStr
		}
	}
	
	// 5. [优化] 获取导演/主创
	var directors []string
	
	// 方法1: 从 header_info 中查找
	doc.Find(".header_info ol.people li").Each(func(i int, s *goquery.Selection) {
		// 获取职位（可能在不同的元素中）
		job := strings.TrimSpace(s.Find("p.character").Text())
		if job == "" {
			job = strings.TrimSpace(s.Find(".character").Text())
		}
		
		// 获取人名
		name := strings.TrimSpace(s.Find("p a").Text())
		if name == "" {
			name = strings.TrimSpace(s.Find("a").First().Text())
		}
		if name == "" {
			name = strings.TrimSpace(s.Find("p").First().Text())
		}
		
		if name == "" {
			return
		}
		
		shouldAdd := false
		if item.Type == "movie" {
			// 电影：查找导演
			if strings.Contains(job, "Director") || 
			   strings.Contains(job, "导演") ||
			   strings.Contains(job, "Directed by") {
				shouldAdd = true
			}
		} else if item.Type == "tv" {
			// 电视剧：查找主创
			jobLower := strings.ToLower(job)
			if strings.Contains(jobLower, "creator") || 
			   strings.Contains(job, "主创") || 
			   strings.Contains(job, "创作") ||
			   strings.Contains(jobLower, "created by") ||
			   strings.Contains(job, "原作") {
				shouldAdd = true
			}
		}
		
		if shouldAdd {
			directors = append(directors, name)
		}
	})
	
	// 方法2: 如果方法1没找到，尝试其他选择器（针对TV）
	if len(directors) == 0 && item.Type == "tv" {
		doc.Find("ol.people.credits li").Each(func(i int, s *goquery.Selection) {
			job := strings.TrimSpace(s.Find("p.character").Text())
			name := strings.TrimSpace(s.Find("p a").Text())
			
			if name != "" {
				jobLower := strings.ToLower(job)
				if strings.Contains(jobLower, "creator") || 
				   strings.Contains(job, "主创") {
					directors = append(directors, name)
				}
			}
		})
	}
	
	// 方法3: 最后的备用方案 - 查找所有 li，匹配关键词
	if len(directors) == 0 && item.Type == "tv" {
		doc.Find("section.panel.top_billed ol li").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			if strings.Contains(text, "主创") || 
			   strings.Contains(strings.ToLower(text), "creator") {
				name := strings.TrimSpace(s.Find("a").First().Text())
				if name != "" {
					directors = append(directors, name)
				}
			}
		})
	}
	
	// 去重并拼接
	directorMap := make(map[string]bool)
	var uniqueDirectors []string
	for _, d := range directors {
		if !directorMap[d] {
			directorMap[d] = true
			uniqueDirectors = append(uniqueDirectors, d)
		}
	}
	item.Director = strings.Join(uniqueDirectors, " / ")
	
	// 6. 获取主演 (Top Billed Cast)
	var actors []string
	doc.Find("ol.people.scroller li.card").Each(func(i int, s *goquery.Selection) {
		if i < 3 { // 取前4个
			actorName := strings.TrimSpace(s.Find("p a").First().Text())
			if actorName == "" {
				actorName = strings.TrimSpace(s.Find("p").First().Text())
			}
			if actorName != "" {
				actors = append(actors, actorName)
			}
		}
	})
	item.Actors = strings.Join(actors, " / ")
	
	if item.Title == "" {
		return FailRespWithMsg(c, Fail, "无法获取信息，可能是ID错误")
	}
	
	return SuccessResp(c, item)
}