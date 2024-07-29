//go:build !prod

package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/handler"
	"github.com/kingwrcy/moments/log"
	"github.com/kingwrcy/moments/middleware"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

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
	do.Provide(injector, db.NewDB)
	do.Provide(injector, newEchoEngine)
	do.Provide(injector, handler.NewBaseHandler)

	myLogger := do.MustInvoke[zerolog.Logger](injector)
	tx := do.MustInvoke[*gorm.DB](injector)

	e := do.MustInvoke[*echo.Echo](injector)
	e.Use(middleware.Auth(injector))
	setupRouter(injector)

	migrateTo3(tx, myLogger)
	myLogger.Info().Msgf("服务端启动成功,监听:%d端口...", cfg.Port)
	err = e.Start(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		myLogger.Fatal().Msgf("服务启动失败,错误原因:%s", err)
	}
}
