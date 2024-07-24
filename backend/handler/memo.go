package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"golang.org/x/net/html"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MemoHandler struct {
	base BaseHandler
}

func NewMemoHandler(injector do.Injector) *MemoHandler {
	return &MemoHandler{do.MustInvoke[BaseHandler](injector)}
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
