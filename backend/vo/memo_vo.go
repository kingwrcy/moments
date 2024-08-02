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
type RemoveImageReq struct {
	Img string `json:"img,omitempty"` //图片路径
}
type Music struct {
	ID     string `json:"id,omitempty"`     //MetingJS的音乐ID
	Server string `json:"server,omitempty"` //音乐的平台
	Type   string `json:"type,omitempty"`   //音乐的类型
	Api    string `json:"api,omitempty"`    //MetingJS的服务端API地址
}

type MemoExt struct {
	Music       Music       `json:"music,omitempty"`       //音乐
	DoubanBook  DoubanBook  `json:"doubanBook,omitempty"`  //豆瓣读书
	DoubanMovie DoubanMovie `json:"doubanMovie,omitempty"` //豆瓣电影
	Video       Video       `json:"video,omitempty"`       //视频
}

type Video struct {
	Type  string `json:"type,omitempty"`  //视频类型,online:在线视频,youtube,bilibili
	Value string `json:"value,omitempty"` //视频地址
}

type SaveMemoReq struct {
	ID              int      `json:"id,omitempty"`              //Memo's ID
	Content         string   `json:"content,omitempty"`         //正文
	Ext             MemoExt  `json:"ext"`                       //扩展
	Pinned          *bool    `json:"pinned,omitempty"`          //是否置顶
	ShowType        *int32   `json:"showType,omitempty"`        //是否公开,1:公开,0:私有
	ExternalFavicon string   `json:"externalFavicon,omitempty"` //外部站点favicon
	ExternalTitle   string   `json:"externalTitle,omitempty"`   //外部站点标题
	ExternalUrl     string   `json:"externalUrl,omitempty"`     //外部站点URL
	Imgs            []string `json:"imgs,omitempty"`            //图片列表,最多9张
	Location        string   `json:"location,omitempty"`        //地理位置
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
