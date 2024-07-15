package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kingwrcy/moments/handler"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

func auth(injector do.Injector) echo.MiddlewareFunc {
	cfg := do.MustInvoke[vo.SysConfig](injector)
	ignores := []string{
		"/user/reg",
		"/user/login",
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			for _, url := range ignores {
				if path == url {
					return next(c)
				}
			}
			tokenStr := c.Request().Header.Get("x-api-token")
			if tokenStr == "" {
				return handler.FailResp(c, handler.TokenMissing)
			}
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JwtKey), nil
			})

			if !token.Valid {
				return handler.FailResp(c, handler.TokenInvalid)
			}

			return next(c)
		}
	}
}
