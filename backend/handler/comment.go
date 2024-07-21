package handler

import (
	"errors"
	"fmt"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type CommentHandler struct {
	base BaseHandler
}

func NewCommentHandler(injector do.Injector) *CommentHandler {
	return &CommentHandler{do.MustInvoke[BaseHandler](injector)}
}

func (c CommentHandler) RemoveComment(ctx echo.Context) error {
	context := ctx.(CustomContext)
	currentUser := context.CurrentUser()
	id, err := strconv.Atoi(ctx.QueryParam("id"))
	if err != nil {
		return FailResp(ctx, ParamError)
	}
	var (
		comment db.Comment
		memo    db.Memo
	)
	if err = c.base.db.First(&comment, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailResp(ctx, ParamError)
	}
	if err = c.base.db.First(&memo, comment.MemoId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailResp(ctx, ParamError)
	}

	if currentUser.Id != memo.UserId && currentUser.Id != 1 {
		return FailRespWithMsg(ctx, Fail, "没有权限")
	}
	if c.base.db.Delete(&comment).RowsAffected != 1 {
		return FailRespWithMsg(ctx, Fail, "删除失败")
	}
	return SuccessResp(ctx, h{})
}

func (c CommentHandler) AddComment(ctx echo.Context) error {
	var (
		req     vo.AddCommentReq
		comment db.Comment
		now     = time.Now()
	)
	err := ctx.Bind(&req)
	if err != nil {
		c.base.log.Error().Msgf("发表评论时参数校验失败,原因:%s", err)
		return FailResp(ctx, ParamError)
	}
	if context, ok := ctx.(CustomContext); ok {
		currentUser := context.CurrentUser()
		comment.Username = currentUser.Username
		comment.Author = fmt.Sprintf("%d", currentUser.Id)
	} else {
		comment.Username = req.Username
	}

	comment.Content = req.Content
	comment.Email = req.Email
	comment.CreatedAt = &now
	comment.UpdatedAt = &now
	comment.ReplyTo = req.ReplyTo
	comment.Website = req.Website
	comment.MemoId = req.MemoID

	if err = c.base.db.Save(&comment).Error; err == nil {
		return SuccessResp(ctx, h{})
	}
	return FailRespWithMsg(ctx, Fail, "发表评论失败")
}
