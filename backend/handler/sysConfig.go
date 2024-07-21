package handler

import (
	"encoding/json"
	"errors"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type SysConfigHandler struct {
	base BaseHandler
}

func NewSysConfigHandler(injector do.Injector) *SysConfigHandler {
	return &SysConfigHandler{do.MustInvoke[BaseHandler](injector)}
}

func (s SysConfigHandler) GetConfig(c echo.Context) error {
	var (
		config db.SysConfig
		result vo.SysConfigVO
	)

	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser == nil || currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "需要先登录")
	}
	if err := s.base.db.First(&config).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return SuccessResp(c, h{})
	}
	err := json.Unmarshal([]byte(config.Content), &result)
	if err != nil {
		return FailRespWithMsg(c, Fail, "读取系统配置异常")
	}
	return SuccessResp(c, result)
}

func (s SysConfigHandler) SaveConfig(c echo.Context) error {
	var (
		config db.SysConfig
		result vo.SysConfigVO
	)
	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser == nil || currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "需要先登录")
	}

	if err := c.Bind(&result); err != nil {
		return FailResp(c, ParamError)
	}

	data, err := json.Marshal(result)
	if err != nil {
		return FailRespWithMsg(c, Fail, "读取系统配置异常")
	}

	if err := s.base.db.First(&config).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		config.Content = string(data)
		if err = s.base.db.Save(&config).Error; err != nil {
			return FailRespWithMsg(c, Fail, "保存系统配置异常")
		}
	} else {
		config.Content = string(data)
		if err = s.base.db.Updates(&config).Error; err != nil {
			return FailRespWithMsg(c, Fail, "保存系统配置异常")
		}
	}
	s.base.db.Table("User").Where("id=?", 1).Update("username", result.AdminUserName)
	return SuccessResp(c, h{})
}
