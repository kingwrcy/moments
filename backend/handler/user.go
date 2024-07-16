package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	base BaseHandler
}

func NewUserHandler(injector do.Injector) *UserHandler {
	return &UserHandler{do.MustInvoke[BaseHandler](injector)}
}

func (u UserHandler) Login(c echo.Context) error {
	var req vo.LoginReq
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}

	var user db.User
	err = u.base.db.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return FailRespWithMsg(c, Fail, "用户不存在或密码不正确")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return FailRespWithMsg(c, Fail, "用户不存在或密码不正确")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"userId":   user.Id,
	})

	tokenString, err := token.SignedString([]byte(u.base.cfg.JwtKey))
	if err != nil {
		u.base.log.Error().Msgf("生成jwt token异常:%s", err)
		return FailRespWithMsg(c, Fail, "登录异常")
	}
	return SuccessResp(c, h{
		"token":    tokenString,
		"username": user.Username,
		"id":       user.Id,
	})
}

func (u UserHandler) Reg(c echo.Context) error {
	var req vo.RegReq
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}

	var user db.User
	err = u.base.db.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return FailRespWithMsg(c, Fail, "用户不存在或密码不正确")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return FailRespWithMsg(c, Fail, "用户不存在或密码不正确")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"userId":   user.Id,
	})

	tokenString, err := token.SignedString([]byte(u.base.cfg.JwtKey))
	if err != nil {
		u.base.log.Error().Msgf("生成jwt token异常:%s", err)
		return FailRespWithMsg(c, Fail, "登录异常")
	}
	return SuccessResp(c, h{
		"token":    tokenString,
		"username": user.Username,
		"id":       user.Id,
	})
}
