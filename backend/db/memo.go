package db

import (
	"time"
)

type Memo struct {
	Id              int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	Content         string     `gorm:"column:content" json:"content,omitempty"`
	Imgs            string     `gorm:"column:imgs" json:"imgs,omitempty"`
	FavCount        int32      `gorm:"column:favCount;default:0;NOT NULL" json:"favCount,omitempty"`
	CommentCount    int32      `gorm:"column:commentCount;default:0;NOT NULL" json:"commentCount,omitempty"`
	UserId          int32      `gorm:"column:userId;NOT NULL" json:"userId,omitempty"`
	CreatedAt       *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
	UpdatedAt       *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`
	Music163Url     string     `gorm:"column:music163Url" json:"music163Url,omitempty"`
	BilibiliUrl     string     `gorm:"column:bilibiliUrl" json:"bilibiliUrl,omitempty"`
	Location        string     `gorm:"column:location" json:"location,omitempty"`
	ExternalUrl     string     `gorm:"column:externalUrl" json:"externalUrl,omitempty"`
	ExternalTitle   string     `gorm:"column:externalTitle" json:"externalTitle,omitempty"`
	ExternalFavicon string     `gorm:"column:externalFavicon;default:/favicon.png;NOT NULL" json:"externalFavicon,omitempty"`
	Pinned          *bool      `gorm:"column:pinned;default:false;NOT NULL" json:"pinned,omitempty"`
	Ext             string     `gorm:"column:ext;default:{};NOT NULL" json:"ext,omitempty"`
	ShowType        *int32     `gorm:"column:showType;default:1;NOT NULL" json:"showType,omitempty"`
	User            *User      `json:"user,omitempty"`
	Comments        []Comment  `json:"comments,omitempty"`
	Tags            string     `json:"tags,omitempty"`
}

func (m *Memo) TableName() string {
	return "Memo"
}
