package log

import (
	"github.com/kingwrcy/moments/vo"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"os"
	"time"
)

func NewLogger(i do.Injector) (zerolog.Logger, error) {
	cfg := do.MustInvoke[*vo.AppConfig](i)

	var level = zerolog.InfoLevel
	err := level.UnmarshalText([]byte(cfg.LogLevel))
	if err != nil {
		level = zerolog.InfoLevel
	} else {
		zerolog.SetGlobalLevel(level)
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}).With().Timestamp().Logger()
	return logger, nil
}
