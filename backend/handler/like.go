package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/pkg/util"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LikeHandler struct {
	base BaseHandler
	hc   http.Client
}

func NewLikeHandler(injector do.Injector) *LikeHandler {
	return &LikeHandler{
		base: do.MustInvoke[BaseHandler](injector),
		hc:   http.Client{},
	}
}

// LikeMemo godoc
//
//	@Tags		Memo
//	@Summary	点赞memo
//	@Accept		json
//	@Produce	json
//	@Param		id	query	int	true	"memoID"
//	@Success	200
//
// @Router /api/like/add [post]
func (l LikeHandler) AddLike(c echo.Context) error {
	var (
		memo        db.Memo
		like        db.Like
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
		token       string
	)

	var currentMemo db.Memo
	if err := l.base.db.Clauses(clause.Locking{Strength: "UPDATE"}).First(&currentMemo, c.QueryParam("id")).Error; err != nil {
		l.base.log.Error().Err(err).Msg("获取memo记录失败")
		return FailRespWithMsg(c, Fail, "系统繁忙，请稍后再试")
	}
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	l.base.db.First(&sysConfig)
	_ = json.Unmarshal([]byte(sysConfig.Content), &sysConfigVO)

	if sysConfigVO.EnableGoogleRecaptcha {
		token = c.QueryParam("token")
		if token == "" {
			return FailRespWithMsg(c, ParamError, "token不能为空")
		}
		if err := checkGoogleRecaptcha(l.base.log, sysConfigVO, token); err != nil {
			return FailRespWithMsg(c, Fail, err.Error())
		}
	}

	if err = l.base.db.First(&memo, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return FailResp(c, ParamError)
	}

	ctx, ok := c.(CustomContext)
	if !ok {
		return FailResp(c, ParamError)
	}
	currentUser := (&ctx).CurrentUser()

	var guestID string
	var guestName string
	if currentUser != nil {
		userId := int(currentUser.Id)
		if err = l.base.db.Where("memoId = ? AND userId = ?", id, userId).First(&like).Error; err == nil {
			return FailRespWithMsg(c, Fail, "您已经点赞过了")
		}
		like = db.Like{
			MemoID: id,
			UserID: &userId,
		}
	} else {
		guestID = c.QueryParam("guestId")
		if !util.IsGuestID(guestID) {
			return FailRespWithMsg(c, ParamError, "无效的访客ID")
		}
		// 检查是否存在关联评论，如果存在则使用评论的用户名，否则使用guestID作为默认名称
		var comment db.Comment
		err := l.base.db.Where("guestId = ?", guestID).Order("createdAt DESC").First(&comment).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				guestName = guestID
			} else {
				return FailRespWithMsg(c, Fail, "查询评论失败")
			}
		} else {
			// 使用评论处理后的用户名
			guestName = comment.Username
		}
		if err = l.base.db.Where("memoId = ? AND guestId = ?", id, guestID).First(&like).Error; err == nil {
			return FailRespWithMsg(c, Fail, "您已经点赞过了")
		}
		like = db.Like{
			MemoID:    id,
			GuestID:   guestID,
			GuestName: guestName,
		}
	}

	tx := l.base.db.Begin()
	if tx.Error != nil {
		l.base.log.Error().Err(tx.Error).Msg("开启事务失败")
		return FailRespWithMsg(c, Fail, "系统繁忙，请稍后再试")
	}

	if err = tx.Create(&like).Error; err != nil {
		tx.Rollback()
		return FailRespWithMsg(c, Fail, "点赞失败")
	}

	if err = tx.Commit().Error; err != nil {
		return FailRespWithMsg(c, Fail, "提交事务失败")
	}

	return SuccessResp(c, h{})
}

// @Router /api/like/get [post]
func (l LikeHandler) GetLike(c echo.Context) error {
	var (
		likes    []db.Like
		users    []db.User
		likeInfo []map[string]interface{}
	)
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	if err = l.base.db.Where("memoId = ?", id).Find(&likes).Error; err != nil {
		return FailRespWithMsg(c, Fail, "获取点赞信息失败")
	}

	userIDs := make([]int, 0)
	for _, like := range likes {
		if like.UserID != nil {
			userIDs = append(userIDs, int(*like.UserID))
		}
	}
	if len(userIDs) > 0 {
		if err = l.base.db.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
			return FailRespWithMsg(c, Fail, "获取用户信息失败")
		}
	}

	userMap := make(map[int]db.User)
	for _, user := range users {
		userMap[int(user.Id)] = user
	}

	for _, like := range likes {
		info := make(map[string]interface{})
		if like.UserID != nil {
			user, ok := userMap[int(*like.UserID)]
			if ok {
				info["id"] = user.Id
				info["name"] = user.Nickname
			}
		} else {
			info["id"] = like.GuestID
			info["name"] = like.GuestName
		}
		likeInfo = append(likeInfo, info)
	}

	result := h{
		"likes": likeInfo,
		"total": len(likes),
	}

	return SuccessResp(c, result)
}

// @Router /api/like/remove [post]
func (l LikeHandler) RemoveLike(c echo.Context) error {
	var (
		like        db.Like
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
		token       string
	)

	var currentMemo db.Memo
	if err := l.base.db.Clauses(clause.Locking{Strength: "UPDATE"}).First(&currentMemo, c.QueryParam("id")).Error; err != nil {
		l.base.log.Error().Err(err).Msg("获取memo记录失败")
		return FailRespWithMsg(c, Fail, "系统繁忙，请稍后再试")
	}
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	l.base.db.First(&sysConfig)
	_ = json.Unmarshal([]byte(sysConfig.Content), &sysConfigVO)

	if sysConfigVO.EnableGoogleRecaptcha {
		token = c.QueryParam("token")
		if token == "" {
			return FailRespWithMsg(c, ParamError, "token不能为空")
		}
		if err := checkGoogleRecaptcha(l.base.log, sysConfigVO, token); err != nil {
			return FailRespWithMsg(c, Fail, err.Error())
		}
	}

	ctx, ok := c.(CustomContext)
	if !ok {
		return FailResp(c, ParamError)
	}
	currentUser := (&ctx).CurrentUser()

	if currentUser != nil {
		userId := int(currentUser.Id)
		if err = l.base.db.Where("memoId = ? AND userId = ?", id, userId).First(&like).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return FailRespWithMsg(c, Fail, "您还没有点赞过")
		}
	} else {
		guestID := c.QueryParam("guestId")
		if err = l.base.db.Where("memoId = ? AND guestId = ?", id, guestID).First(&like).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return FailRespWithMsg(c, Fail, "您还没有点赞过")
		}
	}

	tx := l.base.db.Begin()
	if tx.Error != nil {
		return FailRespWithMsg(c, Fail, "开启事务失败")
	}

	if err = tx.Delete(&like).Error; err != nil {
		tx.Rollback()
		return FailRespWithMsg(c, Fail, "取消点赞失败")
	}

	if err = tx.Commit().Error; err != nil {
		return FailRespWithMsg(c, Fail, "提交事务失败")
	}

	return SuccessResp(c, h{})
}

// @Router /api/like/setGuestId [post]
func (l LikeHandler) SetGuestId(c echo.Context) error {
	cookie, err := c.Cookie("guestInfo")
	if err == nil {
		decodedValue, err := url.QueryUnescape(cookie.Value)
		if err == nil {
			var data vo.GuestInfo
			if err := json.Unmarshal([]byte(decodedValue), &data); err == nil {
				if time.Now().Unix()-data.TimeStamp < 7*24*60*60 {
					return SuccessResp(c, data.GuestId)
				}
			}
		}
	}

	newGuestId := util.GenerateGuestID()
	data := vo.GuestInfo{
		GuestId:   newGuestId,
		TimeStamp: time.Now().Unix(),
	}

	// 将数据编码为 JSON 字符串，并将其转义以安全存储在 Cookie 中
	cookieData, err := json.Marshal(data)
	if err != nil {
		l.base.log.Error().Msgf("Failed to marshal guest data: %v", err)
	}

	encodedValue := url.QueryEscape(string(cookieData))
	secure := c.Scheme() == "https"
	c.SetCookie(&http.Cookie{
		Name:     "guestInfo",
		Value:    encodedValue,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return SuccessResp(c, newGuestId)
}
