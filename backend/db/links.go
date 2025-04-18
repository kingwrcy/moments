package db

import (
	"time"
)

type Link struct {
	Id        int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	Name      string     `gorm:"column:name" json:"name,omitempty"`
	Icon      string     `gorm:"column:icon" json:"icon,omitempty"`
	Url       string     `gorm:"column:url" json:"url,omitempty"`
	Desc      string     `gorm:"column:desc" json:"desc,omitempty"`
	CreatedAt *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`
}

func (n *Link) TableName() string {
	return "Link"
}
