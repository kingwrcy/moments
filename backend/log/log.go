package log

import (
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"os"
	"time"
)

func NewLogger(_ do.Injector) (zerolog.Logger, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}).With().Timestamp().Logger()
	return logger, nil
}
