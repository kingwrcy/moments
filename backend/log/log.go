package log

import (
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"os"
)

func NewLogger(_ do.Injector) (zerolog.Logger, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	return logger, nil
}
