package handler

import (
	"net/url"
	"strconv"
	"time"

	"github.com/kingwrcy/moments/db"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

type LinksHandler struct {
	base BaseHandler
}

func NewLinksHandler(injector do.Injector) *LinksHandler {
	return &LinksHandler{do.MustInvoke[BaseHandler](injector)}
}

// 添加友情链接
// @Router /api/links/add [post]
func (n LinksHandler) AddLinks(c echo.Context) error {
	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser == nil || currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "你没有权限添加友情链接")
	}

	var links db.Links
	if err := c.Bind(&links); err != nil {
		return FailResp(c, ParamError)
	}

	if links.LinksUrl != "" {
		parsedUrl, err := url.Parse(links.LinksUrl)
		if err != nil || (parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https") {
			return FailRespWithMsg(c, Fail, "必须以 http 或 https 开头")
		}
	}

	now := time.Now()
	links.CreatedAt = &now
	links.UpdatedAt = &now

	if err := n.base.db.Create(&links).Error; err != nil {
		return FailRespWithMsg(c, Fail, "添加友情链接失败")
	}

	return SuccessResp(c, links)
}

// 获取友情链接列表
// @Router /api/links/list [post]
func (n LinksHandler) GetLinksList(c echo.Context) error {
	var linkss []db.Links
	if err := n.base.db.Find(&linkss).Error; err != nil {
		return FailRespWithMsg(c, Fail, "获取友情链接列表失败")
	}
	return SuccessResp(c, linkss)
}

// 删除友情链接
// @Router /api/links/delete [post]
func (n LinksHandler) DeleteLinks(c echo.Context) error {
	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser == nil || currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "你没有权限删除友情链接")
	}

	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	if err := n.base.db.Delete(&db.Links{}, id).Error; err != nil {
		return FailRespWithMsg(c, Fail, "删除友情链接失败")
	}

	return SuccessResp[map[string]string](c, map[string]string{"message": "友情链接删除成功"})
}
