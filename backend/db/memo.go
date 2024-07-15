package db

import (
	"database/sql"
	"time"
)

type Memo struct {
	Id              int32          `gorm:"column:id;primary_key;NOT NULL"`
	Content         sql.NullString `gorm:"column:content"`
	Imgs            sql.NullString `gorm:"column:imgs"`
	FavCount        int32          `gorm:"column:favCount;default:0;NOT NULL"`
	CommentCount    int32          `gorm:"column:commentCount;default:0;NOT NULL"`
	UserId          int32          `gorm:"column:userId;NOT NULL"`
	CreatedAt       time.Time      `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt;NOT NULL"`
	Music163Url     sql.NullString `gorm:"column:music163Url"`
	BilibiliUrl     sql.NullString `gorm:"column:bilibiliUrl"`
	Location        sql.NullString `gorm:"column:location"`
	ExternalUrl     sql.NullString `gorm:"column:externalUrl"`
	ExternalTitle   sql.NullString `gorm:"column:externalTitle"`
	ExternalFavicon string         `gorm:"column:externalFavicon;default:/favicon.png;NOT NULL"`
	Pinned          string         `gorm:"column:pinned;default:false;NOT NULL"`
	Ext             string         `gorm:"column:ext;default:{};NOT NULL"`
	ShowType        int32          `gorm:"column:showType;default:1;NOT NULL"`
}

func (m *Memo) TableName() string {
	return "Memo"
}
