package main

import (
	"github.com/google/uuid"
	"github.com/kingwrcy/moments/vo"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strings"
)

func handleEmptyConfig(log zerolog.Logger, cfg *vo.AppConfig) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Error().Msgf("获取当前工作目录异常:%s", err.Error())
		return
	}
	log.Debug().Str("数据库[DB]", cfg.DB).
		Int("端口[PORT]", cfg.Port).
		Str("JWT密钥[JWT_KEY]", cfg.JwtKey).
		Str("上传目录[UPLOAD_DIR]", cfg.UploadDir).
		Str("日志级别[LOG_LEVEL]", cfg.LogLevel).
		Bool("是否启用Swagger文档[ENABLE_SWAGGER]", cfg.EnableSwagger).
		Bool("是否输出SQL[ENABLE_SQL_OUTPUT]", cfg.EnableSQLOutput).
		Msgf("基本信息")
	if cfg.DB != "" && cfg.UploadDir != "" {
		return
	}
	log.Debug().Msgf("没有配置默认所必需的环境变量,在当前目录[%s]生成[数据库文件-db.sqlite]和[上传目录-upload文件夹]...", currentDir)

	if cfg.DB == "" {
		cfg.DB = filepath.Join(currentDir, "db.sqlite")
		if _, err = os.Stat(cfg.DB); err != nil {
			log.Debug().Msgf("当前目录[%s]没有[db.sqlite]数据库文件,自动生成成功", currentDir)
		}
	}
	if cfg.UploadDir == "" {
		cfg.UploadDir = filepath.Join(currentDir, "upload")
		err = os.MkdirAll(cfg.UploadDir, 0755)
		if err != nil {
			log.Fatal().Msgf("创建upload文件夹异常:%s", err.Error())
		} else {
			log.Debug().Msgf("没有配置[上传目录-upload文件夹],在当前目录[%s]生成[upload]文件夹成功", currentDir)
		}
	}
	if cfg.JwtKey == "" {
		cfg.JwtKey = strings.ReplaceAll(uuid.NewString(), "-", "")
		log.Debug().Msgf("JWT_KEY没有配置,随机生成为%s,每次重启服务需要重新登录,配置后则不会", cfg.JwtKey)
	}
}
