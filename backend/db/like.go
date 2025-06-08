package db

import (
	"time"
)

type Like struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	MemoID    int       `gorm:"not null" json:"memo_id"`
	UserID    *int      `gorm:"" json:"user_id"`
	GuestID   string    `gorm:"" json:"guest_id"`
	GuestName string    `gorm:"" json:"guest_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (l *Like) TableName() string {
	return "Like"
}
