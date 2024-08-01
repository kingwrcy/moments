package db

import (
	"context"
	"errors"
	"github.com/glebarez/sqlite"
	"github.com/kingwrcy/moments/vo"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type myLog struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	log                   zerolog.Logger
	cfg                   *vo.AppConfig
}

func newMyLog(injector do.Injector) *myLog {
	return &myLog{
		log:           do.MustInvoke[zerolog.Logger](injector),
		SlowThreshold: time.Second,
		SourceField:   "source",
		cfg:           do.MustInvoke[*vo.AppConfig](injector),
	}
}

func (m myLog) LogMode(level logger.LogLevel) logger.Interface {
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
	if !m.cfg.EnableSQLOutput {
		return
	}
	elapsed := time.Since(begin)
	sql, _ := fc()
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && m.SkipErrRecordNotFound) {
		m.log.Error().Msgf("[GORM] query error,source = %s ,sql = %s,error= %s", utils.FileWithLineNum(), sql, err)
		return
	}

	if m.SlowThreshold != 0 && elapsed > m.SlowThreshold {
		m.log.Warn().Msgf("[GORM] source = %s slow query sql = %s,耗时 = %s", utils.FileWithLineNum(), sql, elapsed)
		return
	}

	m.log.Debug().Msgf("[GORM] query source = %s sql = %s,耗时 = %s", utils.FileWithLineNum(), sql, elapsed)
}

func NewDB(injector do.Injector) (*gorm.DB, error) {
	cfg := do.MustInvoke[*vo.AppConfig](injector)
	log := do.MustInvoke[zerolog.Logger](injector)
	db, err := gorm.Open(sqlite.Open(cfg.DB), &gorm.Config{
		Logger: newMyLog(injector),
	})
	if err != nil {
		log.Fatal().Msgf("无法连接到数据库: %v,路径:%s", err, cfg.DB)
	} else {
		log.Debug().Msgf("连接数据库路径:%s,成功", cfg.DB)
	}

	// 迁移 schema
	err = db.AutoMigrate(&User{}, &Comment{}, &Memo{}, &SysConfig{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
