package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/handler"
	"github.com/kingwrcy/moments/vo"
	"github.com/rs/zerolog"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"strings"
)

func migrateTo3(tx *gorm.DB, log zerolog.Logger) {
	var (
		count int64
		admin db.User
		item  vo.FullSysConfigVO
	)
	tx.Table("SysConfig").Count(&count)
	if count == 0 {
		log.Info().Msg("初始化默认配置...")
		if err := tx.First(&admin).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			admin.Username = "admin"
			admin.Password = "$2a$12$Ruw0XIDW3IuHmD3WXsRTnOUt/0sfqgKWP3wbsqx5sGcCuebWa6X.i"
			admin.Title = "极简朋友圈"
			admin.Slogan = "修道者，逆天而行，注定要一生孤独。"
			admin.Nickname = "admin"
			admin.EnableS3 = "0"
			admin.Favicon = "/favicon.png"
			admin.CoverUrl = "/cover.webp"
			admin.AvatarUrl = "/avatar.webp"
			if err := tx.Save(&admin).Error; err != nil {
				log.Info().Msgf("用户不存在,初始化[admin/a123456]用户... 失败:%s", err)
			} else {
				log.Info().Msg("用户不存在,初始化[admin/a123456]用户... 成功!")
			}
		}
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
		item.EnableGoogleRecaptcha = false
		item.EnableComment = true
		item.MaxCommentLength = 120
		item.MaxCommentLength = 300
		item.CommentOrder = "desc"
		item.TimeFormat = "timeAgo"
		var sysConfig db.SysConfig

		content, err := json.Marshal(&item)
		if err != nil {
			log.Error().Msgf("初始化默认配置执行异常:%s", err)
		}
		sysConfig.Content = string(content)
		if err := tx.Save(&sysConfig).Error; err == nil {
			log.Info().Msg("初始化默认配置执行成功")
		}

		var memos []db.Memo
		tx.Find(&memos)
		for _, memo := range memos {
			log.Info().Msgf("开始迁移memo id:%d", memo.Id)
			var extMap = map[string]interface{}{}
			var ext vo.MemoExt
			err := json.Unmarshal([]byte(memo.Ext), &extMap)
			if err != nil {
				log.Warn().Msgf("memo id:%d ext属性不是标准的json格式 => %s,忽略..", memo.Id, memo.Ext)
				continue
			}
			if value, exist := extMap["videoUrl"]; exist && value != "" {
				ext.Video.Type = "online"
				ext.Video.Value = value.(string)
			}
			if value, exist := extMap["localVideoUrl"]; exist && value != "" {
				ext.Video.Type = "online"
				ext.Video.Value = value.(string)
			}
			if value, exist := extMap["youtubeUrl"]; exist && value != "" {
				ext.Video.Type = "youtube"
				ext.Video.Value = value.(string)
			}
			if memo.BilibiliUrl != "" {
				ext.Video.Type = "bilibili"
				ext.Video.Value = fmt.Sprintf("<iframe src=\"%s\" scrolling=\"no\" border=\"0\" frameborder=\"no\" framespacing=\"0\" allowfullscreen=\"true\"></iframe>", memo.BilibiliUrl)
			}
			if value, exist := extMap["doubanBook"]; exist && value != nil {
				val := gjson.Get(memo.Ext, "doubanBook")
				ext.DoubanBook.Title = val.Get("title").Str
				ext.DoubanBook.Desc = val.Get("desc").Str
				ext.DoubanBook.Image = val.Get("image").Str
				ext.DoubanBook.Author = val.Get("author").Str
				ext.DoubanBook.Isbn = val.Get("isbn").Str
				ext.DoubanBook.Url = val.Get("url").Str
				ext.DoubanBook.Rating = val.Get("rating").Str
				ext.DoubanBook.PubDate = val.Get("pubDate").Str
				ext.DoubanBook.Id = val.Get("id").Str
			}
			if value, exist := extMap["doubanMovie"]; exist && value != nil {
				val := gjson.Get(memo.Ext, "doubanMovie")
				ext.DoubanMovie.Title = val.Get("title").Str
				ext.DoubanMovie.Desc = val.Get("desc").Str
				ext.DoubanMovie.Image = val.Get("image").Str
				ext.DoubanMovie.Director = val.Get("director").Str
				ext.DoubanMovie.ReleaseDate = val.Get("releaseDate").Str
				ext.DoubanMovie.Url = val.Get("url").Str
				ext.DoubanMovie.Rating = val.Get("rating").Str
				ext.DoubanMovie.Actors = val.Get("actors").Str
				ext.DoubanMovie.Id = val.Get("id").Str
			}

			extContent, _ := json.Marshal(ext)
			memo.Ext = string(extContent)
			newTags := ""

			memoContent, tags := handler.FindAndReplaceTags(memo.Content)
			if len(tags) > 0 {
				memo.Content = memoContent
				newTags = strings.Join(tags, ",")
				if newTags != "" {
					newTags = newTags + ","
				}
				memo.Tags = &newTags
			}

			if err = tx.Save(&memo).Error; err != nil {
				log.Info().Msgf("迁移memo id:%d 成功", memo.Id)
			}

		}
	}
}
