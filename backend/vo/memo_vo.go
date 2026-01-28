package vo

import "time"

type ListMemoReq struct {
	Page            int        `json:"page,omitempty"`            //页码,从1开始
	Size            int        `json:"size,omitempty"`            //页大小,默认10
	Tag             string     `json:"tag,omitempty"`             //标签名称,不支持同时多个标签搜索
	Source          string     `json:"source,omitempty"`          //来源,暂无实现
	Username        string     `json:"username,omitempty"`        //用户名
	Start           *time.Time `json:"start,omitempty"`           //开始时间
	End             *time.Time `json:"end,omitempty"`             //结束时间
	ContentContains string     `json:"contentContains,omitempty"` //内容包含
	ShowType        *int       `json:"showType,omitempty"`        //是否是公开的,1:公开,0:私有
	UserId          *int       `json:"userId,omitempty"`          //用户id
}

type Music struct {
	ID     string `json:"id,omitempty"`     //MetingJS的音乐ID
	Server string `json:"server,omitempty"` //音乐的平台
	Type   string `json:"type,omitempty"`   //音乐的类型
	Api    string `json:"api,omitempty"`    //MetingJS的服务端API地址
}

// 确保这里只有一个 MemoExt 定义，并且包含了 SteamGame 和 TmdbItem
type MemoExt struct {
	Music       Music       `json:"music,omitempty"`       //音乐
	DoubanBook  DoubanBook  `json:"doubanBook,omitempty"`  //豆瓣读书
	DoubanMovie DoubanMovie `json:"doubanMovie,omitempty"` //豆瓣电影
	Video       Video       `json:"video,omitempty"`       //视频
	SteamGame   SteamGame   `json:"steamGame,omitempty"`   //Steam游戏 (新增)
	TmdbItem    TmdbItem    `json:"tmdbItem,omitempty"`    //TMDB影视 (新增)
}

type Video struct {
	Type  string `json:"type,omitempty"`  //视频类型,online:在线视频,youtube,bilibili
	Value string `json:"value,omitempty"` //视频地址
}

// 新增 SteamGame 结构体
type SteamGame struct {
	AppId       string `json:"appId,omitempty"`
	Name        string `json:"name,omitempty"`
	Image       string `json:"image,omitempty"` // header.jpg
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Price       string `json:"price,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
}

// 新增 TmdbItem 结构体
type TmdbItem struct {
	Id           string `json:"id,omitempty"`
	Type         string `json:"type,omitempty"` // movie or tv
	Title        string `json:"title,omitempty"`
	Overview     string `json:"overview,omitempty"`
	PosterPath   string `json:"posterPath,omitempty"` // 完整图片URL
	BackdropPath string `json:"backdropPath,omitempty"`
	ReleaseDate  string `json:"releaseDate,omitempty"`
	VoteAverage  string `json:"voteAverage,omitempty"`
	Director     string `json:"director,omitempty"`
	Actors       string `json:"actors,omitempty"`  
	Url          string `json:"url,omitempty"`
}

type SaveMemoReq struct {
	ID              int        `json:"id,omitempty"`              //Memo's ID
	Content         string     `json:"content,omitempty"`         //正文
	Ext             MemoExt    `json:"ext"`                       //扩展
	Pinned          *bool      `json:"pinned,omitempty"`          //是否置顶
	ShowType        *int32     `json:"showType,omitempty"`        //是否公开,1:公开,0:私有
	ExternalFavicon string     `json:"externalFavicon,omitempty"` //外部站点favicon
	ExternalTitle   string     `json:"externalTitle,omitempty"`   //外部站点标题
	ExternalUrl     string     `json:"externalUrl,omitempty"`     //外部站点URL
	Imgs            []string   `json:"imgs,omitempty"`            //图片列表,最多9张
	Location        string     `json:"location,omitempty"`        //地理位置
	Tags            []string   `json:"tags,omitempty"`            //标签数组
	CreatedAt       *time.Time `json:"createdAt,omitempty"`       //创建时间
}

type DoubanMovie struct {
	Id          string `json:"id,omitempty"`          //豆瓣电影ID
	Url         string `json:"url,omitempty"`         //豆瓣电影URL
	Title       string `json:"title,omitempty"`       //标题
	Desc        string `json:"desc,omitempty"`        //描述
	Image       string `json:"image,omitempty"`       //主图
	Director    string `json:"director,omitempty"`    //导演
	Rating      string `json:"rating,omitempty"`      //评分
	ReleaseDate string `json:"releaseDate,omitempty"` //上映日期
	Actors      string `json:"actors,omitempty"`      //演员
	Runtime     string `json:"runtime,omitempty"`     //片长
}

type DoubanBook struct {
	Id       string `json:"id,omitempty"`       //豆瓣读书ID
	Url      string `json:"url,omitempty"`      //豆瓣读书URL
	Title    string `json:"title,omitempty"`    //标题
	Desc     string `json:"desc,omitempty"`     //描述
	Image    string `json:"image,omitempty"`    //主图
	Isbn     string `json:"isbn,omitempty"`     //ISBN
	Author   string `json:"author,omitempty"`   //作者
	Rating   string `json:"rating,omitempty"`   //评分
	PubDate  string `json:"pubDate,omitempty"`  //发布日期
	Keywords string `json:"keywords,omitempty"` //关键字
}

type ImgConfig struct {
	Url      *string `json:"url,omitempty"`
	ThumbUrl *string `json:"thumbUrl,omitempty"`
}