package handler

import (
	"encoding/json"
	"errors"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
	"strconv"
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
		req        vo.ListMemoReq
		list       []db.Memo
		pinnedList []db.Memo
		total      int64
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

	offset := (req.Page - 1) * req.Size

	tx := m.base.db.Preload("User", func(x *gorm.DB) *gorm.DB {
		return x.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl")
	}).Where("pinned = 0")

	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser != nil {
		if req.Source == "" || req.Source == "square" {
			//tx = tx.Where("userId = ?", currentUser.Id)
		}
	}
	tx.Order("createdAt desc").Limit(req.Size).Offset(offset).Find(&list)
	tx.Count(&total)

	m.base.db.Preload("User", func(x *gorm.DB) *gorm.DB {
		return x.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl")
	}).Where("pinned = 1").Find(&pinnedList)

	list = append(pinnedList, list...)

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

func (m MemoHandler) SaveMemo(c echo.Context) error {
	var (
		req     vo.SaveMemoReq
		imgStr  string
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
	memo.Content = req.Content
	for _, img := range req.Imgs {
		imgStr += img + ","
	}

	bytes, _ := json.Marshal(req.Ext)
	extJson = string(bytes)

	memo.Imgs = imgStr
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
