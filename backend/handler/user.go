package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	var (
		req   vo.RegReq
		count int64
		user  db.User
		now   = time.Now()
	)
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}

	if len(req.Username) < 3 {
		return FailRespWithMsg(c, Fail, "用户名最少3个字符")
	}
	if req.Password != req.RepeatPassword {
		return FailRespWithMsg(c, Fail, "两次密码不一致")
	}
	u.base.db.Table("User").Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return FailRespWithMsg(c, Fail, "用户名已存在")
	}
	user.Username = req.Username
	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		u.base.log.Error().Msgf("密码加密异常:%s", err)
		return FailRespWithMsg(c, Fail, "密码加密异常")
	}
	user.Password = string(pwd)
	user.CreatedAt = &now
	user.UpdatedAt = &now
	user.Nickname = req.Username
	user.AvatarUrl = "/avatar.webp"
	user.Slogan = "修道者，逆天而行，注定要一生孤独。"
	user.CoverUrl = "/cover.webp"
	if err := u.base.db.Save(&user).Error; err != nil {
		u.base.log.Error().Msgf("注册用户异常:%s", err)
		return FailRespWithMsg(c, Fail, "注册用户异常")
	}
	return SuccessResp(c, h{})
}

func (u UserHandler) ProfileForUser(c echo.Context) error {
	username := c.Param("username")
	var user db.User
	u.base.db.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl").Find(&user, "username = ?", username)
	return SuccessResp(c, user)
}

func (u UserHandler) Profile(c echo.Context) error {

	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser == nil {
		u.base.db.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl").First(&currentUser)
	}

	return SuccessResp(c, currentUser)
}

func (u UserHandler) SaveProfile(c echo.Context) error {
	var (
		req  vo.ProfileReq
		user db.User
	)
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}
	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser == nil {
		return FailResp(c, TokenMissing)
	}
	u.base.db.Find(&user, currentUser.Id)
	if req.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			return FailResp(c, Fail)
		}
		user.Password = string(password)
	}
	user.Nickname = req.Nickname
	user.AvatarUrl = req.AvatarUrl
	user.Slogan = req.Slogan
	user.CoverUrl = req.CoverUrl

	if err := u.base.db.Save(&user).Error; err != nil {
		return FailResp(c, Fail)
	}
	return SuccessResp(c, h{})
}
