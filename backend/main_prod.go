//go:build prod

package main

import (
	"embed"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kingwrcy/moments/db"
	_ "github.com/kingwrcy/moments/docs"
	"github.com/kingwrcy/moments/handler"
	"github.com/kingwrcy/moments/log"
	myMiddleware "github.com/kingwrcy/moments/middleware"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	_ "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"io/fs"
	"net/http"
)

//go:embed public/*
var staticFiles embed.FS

func newEchoEngine(_ do.Injector) (*echo.Echo, error) {
	e := echo.New()
	return e, nil
}

func main() {

	injector := do.New()
	var cfg vo.AppConfig

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		fmt.Printf("读取配置文件异常:%s", err)
		return
	}

	do.ProvideValue(injector, cfg)
	do.Provide(injector, log.NewLogger)

	myLogger := do.MustInvoke[zerolog.Logger](injector)
	handleEmptyConfig(myLogger, cfg)

	do.Provide(injector, db.NewDB)
	do.Provide(injector, newEchoEngine)
	do.Provide(injector, handler.NewBaseHandler)

	tx := do.MustInvoke[*gorm.DB](injector)

	e := do.MustInvoke[*echo.Echo](injector)
	e.Use(myMiddleware.Auth(injector))

	setupRouter(injector)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Root:       "public", // because files are located in `web` directory in `webAssets` fs
		Filesystem: http.FS(staticFiles),
	}))

	e.FileFS("/*", "public/index.html", staticFiles)

	migrateTo3(tx, myLogger)
	myLogger.Info().Msgf("服务端启动成功,监听:%d端口...", cfg.Port)
	err = e.Start(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		myLogger.Fatal().Msgf("服务启动失败,错误原因:%s", err)
	}
}

func isEmbedFSEmpty(e embed.FS, path string) (bool, error) {
	entries, err := fs.ReadDir(e, path)
	if err != nil {
		return false, err
	}
	return len(entries) == 0, nil
}
