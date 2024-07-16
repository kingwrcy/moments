package db

import (
	"context"
	"github.com/kingwrcy/moments/vo"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type myLog struct {
	log zerolog.Logger
}

func newMyLog(injector do.Injector) *myLog {
	return &myLog{log: do.MustInvoke[zerolog.Logger](injector)}
}

func (m myLog) LogMode(level logger.LogLevel) logger.Interface {
	if level == logger.Warn {
		m.log.Level(zerolog.WarnLevel)
	} else if level == logger.Error {
		m.log.Level(zerolog.ErrorLevel)
	} else {
		m.log.Level(zerolog.InfoLevel)
	}
	return m
}

func (m myLog) Info(ctx context.Context, s string, i ...interface{}) {
	m.log.Info().Msgf(s, i)
}

func (m myLog) Warn(ctx context.Context, s string, i ...interface{}) {
	m.log.Warn().Msgf(s, i)
}

func (m myLog) Error(ctx context.Context, s string, i ...interface{}) {
	m.log.Error().Msgf(s, i)
}

func (m myLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {

}

func NewDB(injector do.Injector) (*gorm.DB, error) {
	cfg := do.MustInvoke[vo.SysConfig](injector)
	db, err := gorm.Open(sqlite.Open(cfg.DB), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 迁移 schema
	err = db.AutoMigrate(&User{}, &Comment{}, &Memo{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
