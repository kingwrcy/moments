package db

import (
	"time"
)

type Like struct {
	Id        int        `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	MemoID    int        `gorm:"column:memoId;NOT NULL" json:"memoId,omitempty"`
	UserID    *int       `gorm:"column:userId" json:"userId,omitempty"`
	GuestID   string     `gorm:"column:guestId" json:"guestId,omitempty"`
	GuestName string     `gorm:"column:guestName" json:"guestName,omitempty"`
	CreatedAt *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
}

func (l *Like) TableName() string {
	return "Like"
}
