package main

import (
	"github.com/kingwrcy/moments/handler"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func setupRouter(injector do.Injector) {
	userHandler := handler.NewUserHandler(injector)
	memoHandler := handler.NewMemoHandler(injector)
	commentHandler := handler.NewCommentHandler(injector)
	sycConfigHandler := handler.NewSysConfigHandler(injector)
	fileHandler := handler.NewFileHandler(injector)
	tagHandler := handler.NewTagHandler(injector)
	e := do.MustInvoke[*echo.Echo](injector)
	cfg := do.MustInvoke[*vo.AppConfig](injector)

	api := e.Group("/api")

	userGroup := api.Group("/user")
	userGroup.POST("/login", userHandler.Login)
	userGroup.POST("/reg", userHandler.Reg)
	userGroup.POST("/profile", userHandler.Profile)
	userGroup.POST("/profile/:username", userHandler.ProfileForUser)
	userGroup.POST("/saveProfile", userHandler.SaveProfile)

	memoGroup := api.Group("/memo")
	memoGroup.POST("/list", memoHandler.ListMemos)
	memoGroup.POST("/save", memoHandler.SaveMemo)
	memoGroup.POST("/remove", memoHandler.RemoveMemo)
	memoGroup.POST("/like", memoHandler.LikeMemo)
	memoGroup.POST("/get", memoHandler.GetMemo)
	memoGroup.POST("/setPinned", memoHandler.SetPinned)
	memoGroup.POST("/getFaviconAndTitle", memoHandler.GetFaviconAndTitle)
	memoGroup.POST("/getDoubanMovieInfo", memoHandler.GetDoubanMovieInfo)
	memoGroup.POST("/getDoubanBookInfo", memoHandler.GetDoubanBookInfo)
	memoGroup.POST("/removeImage", memoHandler.RemoveImage)

	commentGroup := api.Group("/comment")
	commentGroup.POST("/add", commentHandler.AddComment)
	commentGroup.POST("/remove", commentHandler.RemoveComment)

	sycConfigGroup := api.Group("/sysConfig")
	sycConfigGroup.POST("/save", sycConfigHandler.SaveConfig)
	sycConfigGroup.POST("/get", sycConfigHandler.GetConfig)
	sycConfigGroup.POST("/getFull", sycConfigHandler.GetFullConfig)

	tagGroup := api.Group("/tag")
	tagGroup.POST("/list", tagHandler.List)

	e.GET("/upload/:filename", fileHandler.Get)
	e.POST("/api/file/upload", fileHandler.Upload)
	e.POST("/api/file/s3PreSigned", fileHandler.S3PreSigned)

	if cfg.EnableSwagger {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

}
