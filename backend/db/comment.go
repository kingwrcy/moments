package db

import (
	"database/sql"
	"time"
)

type Comment struct {
	Id        int32          `gorm:"column:id;primary_key;NOT NULL"`
	Content   sql.NullString `gorm:"column:content"`
	ReplyTo   sql.NullString `gorm:"column:replyTo"`
	Username  sql.NullString `gorm:"column:username"`
	Email     sql.NullString `gorm:"column:email"`
	Website   sql.NullString `gorm:"column:website"`
	CreatedAt time.Time      `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt time.Time      `gorm:"column:updatedAt;NOT NULL"`
	MemoId    int32          `gorm:"column:memoId;NOT NULL"`
	Author    sql.NullString `gorm:"column:author"`
}

func (c *Comment) TableName() string {
	return "Comment"
}
