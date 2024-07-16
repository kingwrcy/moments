package main

import (
	"github.com/kingwrcy/moments/handler"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

func setupRouter(injector do.Injector) {
	userHandler := handler.NewUserHandler(injector)
	memoHandler := handler.NewMemoHandler(injector)
	e := do.MustInvoke[*echo.Echo](injector)

	api := e.Group("/api")
	userGroup := api.Group("/user")
	userGroup.POST("/login", userHandler.Login)

	memoGroup := api.Group("/memo")
	memoGroup.POST("/list", memoHandler.ListMemos)
}
