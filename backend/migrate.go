package main

import (
	"encoding/json"
	"errors"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func migrateTo3(tx *gorm.DB, log zerolog.Logger) {
	var (
		count int64
		admin db.User
		item  vo.SysConfigVO
	)
	tx.Table("SysConfig").Count(&count)
	if count == 0 {
		if err := tx.First(&admin).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		log.Info().Msg("开始执行迁移...")
		item.AdminUserName = admin.Username
		item.Css = admin.Css
		item.Js = admin.Js
		item.BeiAnNo = admin.BeianNo
		item.Favicon = admin.Favicon
		item.Title = admin.Title
		if admin.EnableS3 == "0" {
			item.EnableS3 = false
		} else {
			item.EnableS3 = true
			item.S3 = vo.S3VO{
				Domain:          admin.Domain,
				Bucket:          admin.Bucket,
				Region:          admin.Region,
				AccessKey:       admin.AccessKey,
				SecretKey:       admin.SecretKey,
				Endpoint:        admin.Endpoint,
				ThumbnailSuffix: admin.ThumbnailSuffix,
			}
		}

		var sysConfig db.SysConfig

		content, err := json.Marshal(&item)
		if err != nil {
			log.Error().Msgf("迁移脚本执行异常:%s", err)
		}
		sysConfig.Content = string(content)
		if err := tx.Save(&sysConfig).Error; err == nil {
			log.Info().Msg("迁移脚本执行成功")
		}
	}
}
