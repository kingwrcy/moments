package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	model "github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/handler"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
	"strings"
)

func Auth(injector do.Injector) echo.MiddlewareFunc {
	cfg := do.MustInvoke[*vo.AppConfig](injector)
	db := do.MustInvoke[*gorm.DB](injector)
	//zlog := do.MustInvoke[zerolog.Logger](injector)
	ignores := []string{
		"/api/user/reg",
		"/api/user/login",
		"/api/memo/list",
		"/api/user/profile",
		"/api/sysConfig/get",
		"/api/memo/like",
		"/api/comment/add",
		"/api/memo/get",
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !strings.HasPrefix(c.Request().URL.Path, "/api") {
				return next(c)
			}
			tokenStr := c.Request().Header.Get("x-api-token")
			cc := handler.CustomContext{Context: c}
			//zlog.Info().Msgf("token :%s", tokenStr)
			if tokenStr != "" {
				token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
					return []byte(cfg.JwtKey), nil
				})

				if err != nil || !token.Valid {
					return handler.FailResp(c, handler.TokenInvalid)
				}

				claims := token.Claims.(jwt.MapClaims)
				//zlog.Info().Msgf("user id :%v", claims["userId"])

				var user model.User
				db.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl").First(&user, claims["userId"])
				cc.SetUser(&user)
				return next(cc)
			} else {
				path := c.Request().URL.Path
				for _, url := range ignores {
					if path == url {
						return next(cc)
					}
				}
				if strings.HasPrefix(path, "/upload") || strings.HasPrefix(path, "/api/user/profile/") {
					return next(cc)
				}
				return handler.FailResp(c, handler.TokenMissing)
			}
		}
	}
}
