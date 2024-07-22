package main

import (
	"github.com/kingwrcy/moments/handler"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

func setupRouter(injector do.Injector) {
	userHandler := handler.NewUserHandler(injector)
	memoHandler := handler.NewMemoHandler(injector)
	commentHandler := handler.NewCommentHandler(injector)
	sycConfigHandler := handler.NewSysConfigHandler(injector)
	fileHandler := handler.NewFileHandler(injector)
	e := do.MustInvoke[*echo.Echo](injector)

	api := e.Group("/api")

	userGroup := api.Group("/user")
	userGroup.POST("/login", userHandler.Login)
	userGroup.POST("/reg", userHandler.Reg)
	userGroup.POST("/profile", userHandler.Profile)
	userGroup.POST("/saveProfile", userHandler.SaveProfile)

	memoGroup := api.Group("/memo")
	memoGroup.POST("/list", memoHandler.ListMemos)
	memoGroup.POST("/save", memoHandler.SaveMemo)
	memoGroup.POST("/remove", memoHandler.RemoveMemo)
	memoGroup.POST("/like", memoHandler.LikeMemo)
	memoGroup.POST("/get", memoHandler.GetMemo)
	memoGroup.POST("/setPinned", memoHandler.SetPinned)

	commentGroup := api.Group("/comment")
	commentGroup.POST("/add", commentHandler.AddComment)
	commentGroup.POST("/remove", commentHandler.RemoveComment)

	sycConfigGroup := api.Group("/sysConfig")
	sycConfigGroup.POST("/save", sycConfigHandler.SaveConfig)
	sycConfigGroup.POST("/get", sycConfigHandler.GetConfig)
	sycConfigGroup.POST("/getFull", sycConfigHandler.GetFullConfig)

	e.GET("/api/file/get/:filename", fileHandler.Get)
	e.POST("/api/file/upload", fileHandler.Upload)
	e.POST("/api/file/s3PreSigned", fileHandler.S3PreSigned)
}
