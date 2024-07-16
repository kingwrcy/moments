package db

import (
	"time"
)

type Comment struct {
	Id        int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	Content   string     `gorm:"column:content" json:"content,omitempty"`
	ReplyTo   string     `gorm:"column:replyTo" json:"replyTo,omitempty"`
	Username  string     `gorm:"column:username" json:"username,omitempty"`
	Email     string     `gorm:"column:email" json:"email,omitempty"`
	Website   string     `gorm:"column:website" json:"website,omitempty"`
	CreatedAt *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`
	MemoId    int32      `gorm:"column:memoId;NOT NULL" json:"memoId,omitempty"`
	Author    string     `gorm:"column:author" json:"author,omitempty"`
	Memo      *Memo      `json:"memo,omitempty"`
}

func (c *Comment) TableName() string {
	return "Comment"
}
