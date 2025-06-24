package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/pkg/mail"
	"github.com/kingwrcy/moments/pkg/util"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type CommentHandler struct {
	base BaseHandler
}

func NewCommentHandler(injector do.Injector) *CommentHandler {
	return &CommentHandler{do.MustInvoke[BaseHandler](injector)}
}

// RemoveComment godoc
//
//	@Tags		Comment
//	@Summary	删除评论
//	@Accept		json
//	@Produce	json
//	@Param		id			query	int		true	"评论ID"
//	@Param		x-api-token	header	string	true	"登录TOKEN"
//	@Success	200
//	@Router		/api/comment/remove [post]
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

		var result map[string]any
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

// AddComment godoc
//
//	@Tags		Comment
//	@Summary	添加评论
//	@Accept		json
//	@Produce	json
//	@Param		object	body	vo.AddCommentReq	true	"添加评论"
//	@Success	200
//
// @Router		/api/comment/add [post]
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

	if !sysConfigVO.EnableComment {
		return FailRespWithMsg(ctx, Fail, "评论未开启")
	}

	if err := checkGoogleRecaptcha(c.base.log, sysConfigVO, req.Token); err != nil {
		return FailRespWithMsg(ctx, Fail, err.Error())
	}
	if context, ok := ctx.(CustomContext); ok {
		currentUser := context.CurrentUser()
		if currentUser == nil {
			if req.GuestID == "" || !util.IsGuestID(req.GuestID) {
				return FailRespWithMsg(ctx, ParamError, "无效的访客ID")
			}
			comment.GuestID = req.GuestID

			// 使用统一的用户名处理逻辑
			isConflict := func(name string) bool {
				var count int64
				c.base.db.Model(&db.User{}).Where("username = ? OR nickname = ?", name, name).Count(&count)
				return count > 0
			}
			comment.Username = util.NormalizeUsername(req.Username, isConflict)
			if comment.Username == "" {
				comment.Username = req.GuestID
			}

			comment.Email = req.Email
			comment.Website = req.Website

			if req.Username != "" {
				// 获取该访客最新的评论姓名
				var latestComment db.Comment
				err := c.base.db.Where("guestId = ?", req.GuestID).
					Order("createdAt DESC").
					First(&latestComment).Error

				// 如果存在最新姓名，则使用该姓名更新所有记录
				if err == nil && latestComment.Username != "" {
					req.Username = latestComment.Username
				}

				// 更新同访客ID的所有评论
				c.base.db.Model(&db.Comment{}).
					Where("guestId = ?", req.GuestID).
					Update("username", req.Username)

				// 更新同访客ID的所有点赞
				c.base.db.Model(&db.Like{}).
					Where("guestId = ?", req.GuestID).
					Update("guestName", req.Username)
			}
		} else {
			comment.Username = currentUser.Nickname
			comment.Email = currentUser.Email
			comment.Author = fmt.Sprintf("%d", currentUser.Id)
		}
	}

	comment.Content = req.Content
	comment.CreatedAt = &now
	comment.UpdatedAt = &now
	comment.ReplyTo = req.ReplyTo
	comment.ReplyEmail = req.ReplyEmail
	comment.MemoId = req.MemoID

	if err = c.base.db.Save(&comment).Error; err == nil {
		go func() {
			frontendHost := fmt.Sprintf("%s://%s", ctx.Scheme(), ctx.Request().Host)
			if err = c.commentEmailNotification(comment, frontendHost); err != nil {
				c.base.log.Error().Msgf("邮件通知失败,原因:%s", err)
			}
		}()
		return SuccessResp(ctx, h{})
	}
	return FailRespWithMsg(ctx, Fail, "发表评论失败")
}

func (c CommentHandler) commentEmailNotification(comment db.Comment, host string) error {
	var (
		memo        db.Memo
		user        db.User
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
	)
	c.base.db.First(&memo, comment.MemoId)
	c.base.db.First(&user, memo.UserId)
	c.base.db.First(&sysConfig)
	_ = json.Unmarshal([]byte(sysConfig.Content), &sysConfigVO)

	// 未开启邮件通知
	if !sysConfigVO.EnableEmail {
		return nil
	}

	// 验证邮箱是否为空
	var targetEmail string
	if comment.ReplyTo != "" { // 回复评论
		targetEmail = comment.ReplyEmail
	} else { // 直接评论
		targetEmail = user.Email
	}
	if targetEmail == "" {
		return nil
	}

	// 获取smtp客户端
	client, err := mail.GetSMTPClient(sysConfigVO.SmtpHost, sysConfigVO.SmtpPort, sysConfigVO.SmtpUsername, sysConfigVO.SmtpPassword)
	if err != nil {
		return err
	}
	defer client.Close()
	c.base.log.Info().Msgf("成功连接到SMTP服务器")

	// 通过模板生成邮件内容
	var poster string
	if comment.ReplyTo != "" { // 回复评论
		poster = comment.ReplyTo
	} else { // 直接评论
		poster = user.Nickname
	}
	data := mail.CommentNotificationEmailData{
		Title:     sysConfigVO.Title,
		Host:      host,
		Poster:    poster,
		Commenter: comment.Username,
		CommentAt: comment.CreatedAt,
		Content:   comment.Content,
		MemoId:    comment.MemoId,
	}
	emailbody, err := mail.GenerateCommentNotificationEmail(data)
	if err != nil {
		return err
	}

	getDomain := func(email string) string {
		index := strings.LastIndex(email, "@")
		domain := strings.ToLower(email[index+1:])
		return domain
	}

	// 附加头部字段
	from := sysConfigVO.SmtpUsername
	to := []string{targetEmail}
	subject := sysConfigVO.Title
	domain := getDomain(sysConfigVO.SmtpUsername)
	email := fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"Date: "+time.Now().Format(time.RFC1123Z)+"\r\n"+
			"Message-ID: <"+time.Now().Format("20060102150405")+"@%s>\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=utf-8\r\n"+
			"\r\n"+
			"%s",
		from, to, subject, domain, emailbody)

	// 发送邮件
	if err := client.SendMail(from, to, strings.NewReader(email)); err != nil {
		return err
	}

	c.base.log.Info().Msgf("成功发送邮件")
	return nil
}
