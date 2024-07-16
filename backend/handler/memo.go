package handler

import (
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type MemoHandler struct {
	base BaseHandler
}

func NewMemoHandler(injector do.Injector) *MemoHandler {
	return &MemoHandler{do.MustInvoke[BaseHandler](injector)}
}

func (m MemoHandler) ListMemos(c echo.Context) error {
	var (
		req   vo.ListMemoReq
		list  []db.Memo
		total int64
	)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	tx := m.base.db.Preload("User", func(x *gorm.DB) *gorm.DB {
		return x.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl")
	}).Where("pinned = 0").Find(&list)

	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser != nil {
		if req.Source == "" || req.Source == "square" {
			tx = tx.Where("userId = ?", currentUser.Id)
		}
	}
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}
	offset := (req.Page - 1) * req.Size
	tx.Order("createdAt desc").Limit(req.Size).Offset(offset)

	tx.Count(&total)

	for i, memo := range list {
		var comments []db.Comment
		m.base.db.Where("memoId = ?", memo.Id).Order("createdAt DESC").Limit(5).Find(&comments)
		list[i].Comments = comments
	}

	return SuccessResp(c, h{
		"list":  list,
		"total": total,
	})
}
