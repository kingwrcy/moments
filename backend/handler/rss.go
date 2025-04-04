package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gorilla/feeds"
	"github.com/microcosm-cc/bluemonday"
)

type RssHandler struct {
	base BaseHandler
	hc   http.Client
}

func NewRssHandler(injector do.Injector) *RssHandler {
	return &RssHandler{
		base: do.MustInvoke[BaseHandler](injector),
		hc:   http.Client{},
	}
}

func (r RssHandler) GetRss(c echo.Context) error {
	frontendHost := fmt.Sprintf("%s://%s", c.Scheme(), c.Request().Host)
	rss, err := r.generateRss(frontendHost)
	if err != nil {
		return FailRespWithMsg(c, Fail, "RSS生成失败")
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/rss+xml; charset=utf-8")

	return c.String(http.StatusOK, rss)
}

func (r RssHandler) generateRss(host string) (string, error) {
	var (
		memos       []db.Memo
		user        db.User
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
	)

	// 获取系统设置
	r.base.db.First(&sysConfig)
	_ = json.Unmarshal([]byte(sysConfig.Content), &sysConfigVO)

	// 获取管理员信息
	r.base.db.First(&user, "Username = ?", "admin")

	// 使用自定义RSS
	if sysConfigVO.Rss != "" {
		return "", nil
	}

	// 查询动态
	tx := r.base.db.Preload("User", func(x *gorm.DB) *gorm.DB {
		return x.Select("username", "nickname", "id", "email")
	}).Where("showType = 1")
	tx.Order("createdAt desc").Limit(15).Find(&memos)

	for i := range memos {
		if *memos[i].Pinned {
			memos[i].Content = "**【置顶】**\n" + memos[i].Content
		}
	}

	feed := generateFeed(memos, &sysConfigVO, &user, host)

	return feed.ToRss()

	// // 将RSS内容写入/rss/default_rss.xml
	// target := "/rss/default_rss.xml"
	// dir := filepath.Dir(target)
	// if err := os.MkdirAll(dir, os.ModePerm); err != nil {
	// 	return "", fmt.Errorf("创建目录失败: %w", err)
	// }
	// if err := os.WriteFile(target, []byte(rss), 0644); err != nil {
	// 	return "", fmt.Errorf("写入RSS失败: %w", err)
	// }
}

func generateFeed(memos []db.Memo, sysConfigVO *vo.FullSysConfigVO, user *db.User, host string) *feeds.Feed {
    now := time.Now()
    feed := &feeds.Feed{
        Title:       sysConfigVO.Title,
        Link:        &feeds.Link{Href: fmt.Sprintf("%s/rss", host)},
        Description: user.Slogan,
        Author:      &feeds.Author{Name: user.Nickname, Email: user.Email},
        Created:     now,
    }

    feed.Items = []*feeds.Item{}
    // 定义标题截取的长度
    const maxTitleLength = 20
    for _, memo := range memos {
        memoLink := fmt.Sprintf("%s/memo/%d", host, memo.Id)
        title := ""
        // 检查内容是否为空
        if memo.Content != "" {
            // 按换行符分割内容，取第一行
            lines := strings.Split(memo.Content, "\n")
            if len(lines) > 0 {
                title = lines[0]
                runeTitle := []rune(title)
                if len(runeTitle) > maxTitleLength {
                    title = string(runeTitle[:maxTitleLength]) + "..."
                }
            }
        } else {
            // 若内容为空，使用默认标题格式
            title = fmt.Sprintf("Memo #%d", memo.Id)
        }

        // 获取并格式化标签
        tagStr := ""
        if memo.Tags != nil && *memo.Tags != "" {
            tags := strings.Split(strings.TrimSuffix(*memo.Tags, ","), ",")
            tagStr = fmt.Sprintf("[%s] ", strings.Join(tags, "]["))
        }

        // 将标签添加到标题前面
        title = tagStr + title

        feed.Items = append(feed.Items, &feeds.Item{
            Id:          memoLink,
            Title:       title,
            Link:        &feeds.Link{Href: memoLink},
            Description: parseMarkdownToHtml(getContentWithExt(memo, host)),
            Author:      &feeds.Author{Name: memo.User.Nickname, Email: memo.User.Email},
            Created:     *memo.CreatedAt,
            Updated:     *memo.UpdatedAt,
        })
    }
    return feed
}

func parseMarkdownToHtml(md string) string {
	// 启用扩展
	extensions := parser.NoIntraEmphasis | // 忽略单词内部的强调标记
		parser.Tables | // 解析表格语法
		parser.FencedCode | // 解析围栏代码块
		parser.Strikethrough | // 支持删除线语法
		parser.HardLineBreak | // 将换行符（\n）转换为 <br> 标签
		parser.Footnotes | // 支持脚注语法
		parser.MathJax | // 支持 MathJax 数学公式语法
		parser.SuperSubscript | // 支持上标和下标语法
		parser.EmptyLinesBreakList // 允许两个空行中断列表
	p := parser.NewWithExtensions(extensions)

	// 将 Markdown 解析为 HTML
	html := markdown.ToHTML([]byte(md), p, nil)

	// 清理 HTML（防止 XSS 攻击）
	cleanHTML := bluemonday.UGCPolicy().SanitizeBytes(html)

	return string(cleanHTML)
}

func getContentWithExt(memo db.Memo, host string) string {
	content := memo.Content

	// 处理链接
	if memo.ExternalUrl != "" {
		content += fmt.Sprintf("\n\n[%s](%s)", memo.ExternalTitle, memo.ExternalUrl)
	}

	// 处理图片
	if memo.Imgs != "" {
		imgs := strings.Split(memo.Imgs, ",")
		for _, img := range imgs {
			if img == "" {
				continue
			}

			if strings.HasPrefix(img, "/upload/") {
				img = host + img
			}

			content += fmt.Sprintf("\n\n![%s](%s)", img, img)
		}
	}

	var ext vo.MemoExt
	err := json.Unmarshal([]byte(memo.Ext), &ext)
	if err != nil {
		ext = vo.MemoExt{
			Music:       vo.Music{},
			Video:       vo.Video{},
			DoubanBook:  vo.DoubanBook{},
			DoubanMovie: vo.DoubanMovie{},
		}
	}

	// 处理音乐
	if ext.Music.Server != "" {
		var title, url string
		switch ext.Music.Server {
		// 网易云音乐
		case "netease":
			title = "网易云音乐"
			switch ext.Music.Type {
			case "search":
				ext.Music.ID = "/m/?s=" + ext.Music.ID
			default:
				ext.Music.ID = "?id=" + ext.Music.ID
			}
			url = fmt.Sprintf("https://music.163.com/#/%s%s",
				ext.Music.Type, ext.Music.ID)
		// QQ音乐
		case "tencent":
			title = "QQ音乐"
			switch ext.Music.Type {
			case "song":
				url = fmt.Sprintf("https://y.qq.com/n/ryqq/songDetail/%s", ext.Music.ID)
			case "playlist":
				url = fmt.Sprintf("https://y.qq.com/n/ryqq/playlist/%s", ext.Music.ID)
			case "album":
				url = fmt.Sprintf("https://y.qq.com/n/ryqq/albumDetail/%s", ext.Music.ID)
			case "search":
				url = fmt.Sprintf("https://y.qq.com/n/ryqq/search?w=%s&t=song", ext.Music.ID)
			case "artist":
				url = fmt.Sprintf("https://y.qq.com/n/ryqq/singer/%s", ext.Music.ID)
			default:
			}
		// 酷狗音乐
		case "kugou":
			title = "酷狗音乐"
			switch ext.Music.Type {
			case "song":
				url = fmt.Sprintf("https://www.kugou.com/mixsong/%s.html", ext.Music.ID)
			case "playlist":
				url = fmt.Sprintf("https://www.kugou.com/songlist/%s/", ext.Music.ID)
			case "album":
				url = fmt.Sprintf("https://www.kugou.com/album/info/%s/", ext.Music.ID)
			case "search":
				url = fmt.Sprintf("https://www.kugou.com/yy/html/search.html#searchType=song&searchKeyWord=%s", ext.Music.ID)
			case "artist":
				url = fmt.Sprintf("https://www.kugou.com/singer/info/%s/", ext.Music.ID)
			default:
			}
		// 虾米音乐 已停止服务
		case "xiami":
		// 百度音乐 不可用
		case "baidu":
		default:
		}

		if url != "" {
			content += fmt.Sprintf("\n\n[%s](%s)", title, url)
		}
	}

	// 处理视频
	if ext.Video.Type != "" {
		var title, url string
		switch ext.Video.Type {
		case "online":
			title = "在线视频"
		case "bilibili":
			title = "Bilibili视频"
		case "youtube":
			title = "Youtube视频"
		}
		url = ext.Video.Value
		if ext.Video.Type == "online" && url[:7] == "/upload" {
			url = host + url
		}
		content += fmt.Sprintf("\n\n[%s](%s)", title, url)
	}

	// 处理豆瓣
	if ext.DoubanBook.Url != "" {
		content += fmt.Sprintf("\n\n[%s](%s)", ext.DoubanBook.Title, ext.DoubanBook.Url)
	}
	if ext.DoubanMovie.Url != "" {
		content += fmt.Sprintf("\n\n[%s](%s)", ext.DoubanMovie.Title, ext.DoubanMovie.Url)
	}

	return content
}
