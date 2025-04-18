package handler

import (
	"net/url"
	"strconv"

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

	var link db.Link
	if err := c.Bind(&link); err != nil {
		return FailResp(c, ParamError)
	}

	if link.Url != "" {
		parsedUrl, err := url.Parse(link.Url)
		if err != nil || (parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https") {
			return FailRespWithMsg(c, Fail, "必须以 http 或 https 开头")
		}
	}

	if err := n.base.db.Create(&link).Error; err != nil {
		return FailRespWithMsg(c, Fail, "添加友情链接失败")
	}

	return SuccessResp(c, link)
}

// 获取友情链接列表
// @Router /api/links/list [post]
func (n LinksHandler) GetLinksList(c echo.Context) error {
	var links []db.Link
	if err := n.base.db.Find(&links).Error; err != nil {
		return FailRespWithMsg(c, Fail, "获取友情链接列表失败")
	}
	return SuccessResp(c, links)
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

	if err := n.base.db.Delete(&db.Link{}, id).Error; err != nil {
		return FailRespWithMsg(c, Fail, "删除友情链接失败")
	}

	return SuccessResp[map[string]string](c, map[string]string{"message": "友情链接删除成功"})
}
