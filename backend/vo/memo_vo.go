package vo

import "time"

type ListMemoReq struct {
	Page            int        `json:"page,omitempty"`
	Size            int        `json:"size,omitempty"`
	Tag             string     `json:"tag,omitempty"`
	Source          string     `json:"source,omitempty"`
	Username        string     `json:"username,omitempty"`
	Start           *time.Time `json:"start,omitempty"`
	End             *time.Time `json:"end,omitempty"`
	ContentContains string     `json:"contentContains,omitempty"`
	ShowType        *int       `json:"showType,omitempty"`
}
type RemoveImageReq struct {
	Img string `json:"img,omitempty"`
}
type Music struct {
	ID     string `json:"id,omitempty"`
	Server string `json:"server,omitempty"`
	Type   string `json:"type,omitempty"`
	Api    string `json:"api,omitempty"`
}

type MemoExt struct {
	Music       Music       `json:"music,omitempty"`
	DoubanBook  DoubanBook  `json:"doubanBook,omitempty"`
	DoubanMovie DoubanMovie `json:"doubanMovie,omitempty"`
	Video       Video       `json:"video,omitempty"`
}

type Video struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type SaveMemoReq struct {
	ID              int      `json:"id,omitempty"`
	Content         string   `json:"content,omitempty"`
	Ext             MemoExt  `json:"ext"`
	Pinned          *bool    `json:"pinned,omitempty"`
	ShowType        *int32   `json:"showType,omitempty"`
	ExternalFavicon string   `json:"externalFavicon,omitempty"`
	ExternalTitle   string   `json:"externalTitle,omitempty"`
	ExternalUrl     string   `json:"externalUrl,omitempty"`
	Imgs            []string `json:"imgs,omitempty"`
	Location        string   `json:"location,omitempty"`
}

type DoubanMovie struct {
	Id          string `json:"id,omitempty"`
	Url         string `json:"url,omitempty"`
	Title       string `json:"title,omitempty"`
	Desc        string `json:"desc,omitempty"`
	Image       string `json:"image,omitempty"`
	Director    string `json:"director,omitempty"`
	Rating      string `json:"rating,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	Actors      string `json:"actors,omitempty"`
	Runtime     string `json:"runtime,omitempty"`
}

type DoubanBook struct {
	Id       string `json:"id,omitempty"`
	Url      string `json:"url,omitempty"`
	Title    string `json:"title,omitempty"`
	Desc     string `json:"desc,omitempty"`
	Image    string `json:"image,omitempty"`
	Isbn     string `json:"isbn,omitempty"`
	Author   string `json:"author,omitempty"`
	Rating   string `json:"rating,omitempty"`
	PubDate  string `json:"pubDate,omitempty"`
	Keywords string `json:"keywords,omitempty"`
}
