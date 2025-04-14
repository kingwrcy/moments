package db

import (
	"time"
)

type Links struct {
	Id        int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	LinksName string     `gorm:"column:linksName" json:"linksName,omitempty"`
	LinksIcon string     `gorm:"column:linksIcon" json:"linksIcon,omitempty"`
	LinksUrl  string     `gorm:"column:linksUrl" json:"linksUrl,omitempty"`
	LinksDesc string     `gorm:"column:linksDesc" json:"linksDesc,omitempty"`
	CreatedAt *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`
}

func (n *Links) TableName() string {
	return "Links"
}
