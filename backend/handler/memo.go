package handler

import (
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

type MemoHandler struct {
	base BaseHandler
}

func NewMemoHandler(injector do.Injector) *MemoHandler {
	return &MemoHandler{do.MustInvoke[BaseHandler](injector)}
}

func (m MemoHandler) ListMemos(c echo.Context) error {
	var list []db.Memo
	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser != nil {

	}
	var req vo.ListMemoReq
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}
	offset := (req.Page - 1) * req.Size
	m.base.db.Where("pinned = false", &list).Order("createdAt desc").Limit(req.Size).Offset(offset)
	return SuccessResp(c, h{
		"list": list,
	})
}
