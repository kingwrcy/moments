package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"golang.org/x/net/html"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func (m MemoHandler) ListMemos(c echo.Context) error {
	var (
		req         vo.ListMemoReq
		list        []db.Memo
		pinnedList  []db.Memo
		total       int64
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
	)
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
	totalCondition := m.base.db.Table("Memo").Where("pinned = 0")

	tx := m.base.db.Preload("User", func(x *gorm.DB) *gorm.DB {
		return x.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl")
	}).Where("pinned = 0")

	if req.Start != nil {
		tx = tx.Where("createdAt >= ?", req.Start)
		totalCondition = totalCondition.Where("createdAt >= ?", req.Start)
	}
	if req.End != nil {
		tx = tx.Where("createdAt <= ?", req.End)
		totalCondition = totalCondition.Where("createdAt <= ?", req.End)
	}
	if req.ContentContains != "" {
		tx = tx.Where("content like ?", "%"+req.ContentContains+"%")
		totalCondition = totalCondition.Where("content like ?", "%"+req.ContentContains+"%")
	}
	if req.ShowType != nil && *req.ShowType >= 0 {
		tx = tx.Where("showType=?", req.ShowType)
		totalCondition = totalCondition.Where("showType=?", req.ShowType)
	}
	if req.Tag != "" {
		if strings.Contains(req.Tag, ",") {
			tags := strings.Split(req.Tag, ",")
			for _, tag := range tags {
				tx = tx.Where("tags like ?", fmt.Sprintf("%%%s,%%", tag))
				totalCondition = totalCondition.Where("tags like ?", fmt.Sprintf("%%%s,%%", tag))
			}
		} else {
			tx = tx.Where("tags like ?", fmt.Sprintf("%%%s,%%", req.Tag))
			totalCondition = totalCondition.Where("tags like ?", fmt.Sprintf("%%%s,%%", req.Tag))
		}
	}
	if req.Username != "" {
		var target db.User
		if err := m.base.db.Where("username = ?", req.Username).First(&target).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return FailRespWithMsg(c, Fail, "不存在的用户")
		}
		tx = tx.Where("userId = ?", target.Id)
		totalCondition = totalCondition.Where("userId = ?", target.Id)
	}
	tx.Order("createdAt desc").Limit(req.Size).Offset(offset).Find(&list)
	totalCondition.Count(&total)

	m.base.db.Preload("User", func(x *gorm.DB) *gorm.DB {
		return x.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl")
	}).Where("pinned = 1").Find(&pinnedList)

	list = append(pinnedList, list...)

	for i, memo := range list {
		var comments []db.Comment
		m.base.db.Where("memoId = ?", memo.Id).Order(fmt.Sprintf("createdAt %s", sysConfigVO.CommentOrder)).Limit(5).Find(&comments)
		list[i].Comments = comments
	}

	return SuccessResp(c, h{
		"list":    list,
		"total":   total,
		"hasNext": int64(req.Page*req.Size) < total,
	})
}

func (m MemoHandler) RemoveMemo(c echo.Context) error {
	context := c.(CustomContext)
	currentUser := context.CurrentUser()
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

func (m MemoHandler) LikeMemo(c echo.Context) error {
	var (
		memo db.Memo
	)
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return FailResp(c, ParamError)
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

func parseTags(content string) (string, []string) {
	var tags []string

	// Regular expression to match tags
	tagRegex := regexp.MustCompile(`#(\S+)(\s|$)`)
	// Regular expression to match Markdown headers
	headerRegex := regexp.MustCompile(`(?m)^(#+\s.*)`)

	// Function to check if a line is a Markdown header
	isHeader := func(line string) bool {
		return headerRegex.MatchString(line)
	}

	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if isHeader(line) {
			continue
		}
		// Find all tags in the line
		matches := tagRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				tags = append(tags, match[1])
			}
		}
		// Replace tags in the line
		lines[i] = tagRegex.ReplaceAllString(line, "")
	}

	// Join the lines back together
	processedContent := strings.Join(lines, "\n")
	return processedContent, tags
}

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
	context := c.(CustomContext)
	currentUser := context.CurrentUser()

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

	content, tags := parseTags(req.Content)
	memo.Tags = strings.Join(tags, ",")
	if memo.Tags != "" {
		memo.Tags = memo.Tags + ","
	}
	memo.Content = strings.TrimSpace(content)

	bytes, _ := json.Marshal(req.Ext)
	extJson = string(bytes)

	memo.Imgs = strings.Join(req.Imgs, ",")
	memo.Music163Url = req.Music163Url
	memo.BilibiliUrl = req.BilibiliUrl
	memo.Location = req.Location
	memo.ExternalUrl = req.ExternalUrl
	memo.ExternalTitle = req.ExternalTitle
	memo.ExternalFavicon = req.ExternalFavicon
	memo.Pinned = req.Pinned
	memo.Ext = extJson
	memo.ShowType = int32(req.ShowType)

	if req.ID > 0 {
		m.base.db.Updates(&memo)
	} else {
		m.base.db.Save(&memo)
	}

	return SuccessResp(c, h{})
}

func (m MemoHandler) GetMemo(c echo.Context) error {
	var (
		memo db.Memo
	)
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

	var comments []db.Comment
	tx := m.base.db.Where("memoId = ?", memo.Id).Order("createdAt DESC")
	if latest != "" {
		tx.Limit(5)
	}
	tx.Find(&comments)

	memo.Comments = comments

	return SuccessResp(c, memo)
}

func (m MemoHandler) SetPinned(c echo.Context) error {
	context := c.(CustomContext)
	currentUser := context.CurrentUser()
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

func (m MemoHandler) GetFaviconAndTitle(c echo.Context) error {
	websiteURL := c.QueryParam("url")
	favicon, title, err := getFaviconAndTitle(websiteURL)
	if err != nil {
		return FailResp(c, ParamError)
	}
	return SuccessResp(c, h{
		"favicon": favicon,
		"title":   title,
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
	target := fmt.Sprintf("https://book.douban.com/subject/%s/", id)
	// Request the HTML page.

	req, _ := http.NewRequest("GET", target, nil)
	req.Header.Set("User-Agent", userAgent)
	start := time.Now()
	res, err := m.hc.Do(req)
	m.base.log.Info().Str("豆瓣读书ID", id).Str("耗时", fmt.Sprintf("%f秒", time.Since(start).Seconds()))
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
		req, _ := http.NewRequest("GET", target, nil)
		req.Header.Set("User-Agent", userAgent)
		response, err := m.hc.Do(req)
		if err != nil {
			return FailRespWithMsg(c, Fail, fmt.Sprintf("下载豆瓣电影图片异常:%s", err.Error()))
		}
		defer response.Body.Close()
		client := s3.NewFromConfig(cfg)
		key := fmt.Sprintf("moments/%s/%s", time.Now().Format("2006/01/02"), strings.ReplaceAll(uuid.NewString(), "-", ""))
		_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(sysConfigVo.S3.Bucket),
			Key:    aws.String(key),
			Body:   response.Body,
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

func downloadImage(src string, log zerolog.Logger, conf vo.AppConfig) (string, error) {
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
	return fmt.Sprintf("/api/file/get/%s.jpg", key), err
}

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
		req, _ = http.NewRequest("GET", target, nil)
		req.Header.Set("User-Agent", userAgent)
		response, err := client.Do(req)
		if err != nil {
			return FailRespWithMsg(c, Fail, fmt.Sprintf("下载豆瓣图片异常:%s", err.Error()))
		}
		defer response.Body.Close()
		s3Client := s3.NewFromConfig(cfg)
		key := fmt.Sprintf("moments/%s/%s", time.Now().Format("2006/01/02"), strings.ReplaceAll(uuid.NewString(), "-", ""))
		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(sysConfigVo.S3.Bucket),
			Key:    aws.String(key),
			Body:   response.Body,
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
