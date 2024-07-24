package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
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

func checkGoogleRecaptcha(logger zerolog.Logger, sysConfigVO vo.FullSysConfigVO, token string) error {
	if sysConfigVO.EnableGoogleRecaptcha {
		if token == "" {
			return errors.New("token必填")
		}
		params := url.Values{}
		params.Set("secret", sysConfigVO.GoogleSecretKey)
		params.Set("response", token)

		response, err := http.Post("https://recaptcha.net/recaptcha/api/siteverify?"+params.Encode(), "", nil)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return errors.New("google验证服务无法正常返回")
		}
		resp, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		logger.Info().Str("Action", "评论").Msgf("google resp: %s", resp)

		var result map[string]interface{}
		err = json.Unmarshal(resp, &result)
		if err != nil {
			return err
		}
		if success, ok := result["success"].(bool); ok {
			if success {
				if score, ok := result["score"].(float64); ok {
					if score > 0.5 {
						return nil
					}
				}
			}
		}
		return errors.New("人机校验不通过")
	}
	return nil
}

func (c CommentHandler) AddComment(ctx echo.Context) error {
	var (
		req         vo.AddCommentReq
		comment     db.Comment
		now         = time.Now()
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
	)
	err := ctx.Bind(&req)
	if err != nil {
		c.base.log.Error().Msgf("发表评论时参数校验失败,原因:%s", err)
		return FailResp(ctx, ParamError)
	}
	c.base.db.First(&sysConfig)
	_ = json.Unmarshal([]byte(sysConfig.Content), &sysConfigVO)
	if err := checkGoogleRecaptcha(c.base.log, sysConfigVO, req.Token); err != nil {
		return FailRespWithMsg(ctx, Fail, err.Error())
	}
	if context, ok := ctx.(CustomContext); ok {
		currentUser := context.CurrentUser()
		if currentUser == nil {
			comment.Username = req.Username
		} else {
			comment.Username = currentUser.Username
			comment.Author = fmt.Sprintf("%d", currentUser.Id)
		}
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
