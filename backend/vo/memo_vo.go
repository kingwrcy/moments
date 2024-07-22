package vo

type ListMemoReq struct {
	Page     int    `json:"page,omitempty"`
	Size     int    `json:"size,omitempty"`
	Tag      string `json:"tag,omitempty"`
	Source   string `json:"source,omitempty"`
	Username string `json:"username,omitempty"`
}

type MemoExt struct {
	YoutubeUrl    string `json:"youtube_url,omitempty"`
	VideoUrl      string `json:"video_url,omitempty"`
	LocalVideoUrl string `json:"local_video_url,omitempty"`
}

type SaveMemoReq struct {
	ID              int      `json:"id,omitempty"`
	Content         string   `json:"content,omitempty"`
	Ext             MemoExt  `json:"ext"`
	Pinned          *bool    `json:"pinned,omitempty"`
	ShowType        int      `json:"show_type,omitempty"`
	ExternalFavicon string   `json:"external_favicon,omitempty"`
	ExternalTitle   string   `json:"external_title,omitempty"`
	ExternalUrl     string   `json:"external_url,omitempty"`
	Imgs            []string `json:"imgs,omitempty"`
	Music163Url     string   `json:"music163_url,omitempty"`
	BilibiliUrl     string   `json:"bilibili_url,omitempty"`
	Location        string   `json:"location,omitempty"`
}
