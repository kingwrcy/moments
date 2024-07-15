package main

import (
	"github.com/kingwrcy/moments/handler"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

func setupRouter(injector do.Injector) {
	userHandler := handler.NewUserHandler(injector)
	e := do.MustInvoke[*echo.Echo](injector)

	userGroup := e.Group("/user")
	userGroup.POST("/login", userHandler.Login)
}
