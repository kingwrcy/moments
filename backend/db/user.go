package db

import (
	"database/sql"
	"time"
)

type User struct {
	Id              int32          `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	Username        string         `gorm:"column:username;NOT NULL;unique"  json:"username,omitempty"`
	Nickname        sql.NullString `gorm:"column:nickname" json:"nickname"`
	Password        string         `gorm:"column:password;NOT NULL" json:"password,omitempty"`
	AvatarUrl       sql.NullString `gorm:"column:avatarUrl" json:"avatarUrl"`
	Slogan          sql.NullString `gorm:"column:slogan" json:"slogan"`
	CoverUrl        sql.NullString `gorm:"column:coverUrl" json:"coverUrl"`
	CreatedAt       time.Time      `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt;NOT NULL" json:"updatedAt"`
	EnableS3        string         `gorm:"column:enableS3;default:false;NOT NULL" json:"enableS3,omitempty"`
	Domain          sql.NullString `gorm:"column:domain" json:"domain"`
	Bucket          sql.NullString `gorm:"column:bucket" json:"bucket"`
	Region          sql.NullString `gorm:"column:region" json:"region"`
	AccessKey       sql.NullString `gorm:"column:accessKey" json:"accessKey"`
	SecretKey       sql.NullString `gorm:"column:secretKey" json:"secretKey"`
	Endpoint        sql.NullString `gorm:"column:endpoint" json:"endpoint"`
	ThumbnailSuffix sql.NullString `gorm:"column:thumbnailSuffix" json:"thumbnailSuffix"`
	Favicon         sql.NullString `gorm:"column:favicon" json:"favicon"`
	Title           string         `gorm:"column:title;default:'极简朋友圈';NOT NULL" json:"title,omitempty"`
	BeianNo         sql.NullString `gorm:"column:beianNo" json:"beianNo"`
	Css             sql.NullString `gorm:"column:css" json:"css"`
	Js              sql.NullString `gorm:"column:js" json:"js"`
}

func (u *User) TableName() string {
	return "User"
}
