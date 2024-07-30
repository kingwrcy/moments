package db

import (
	"time"
)

type User struct {
	Id              int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`                             //用户ID
	Username        string     `gorm:"column:username;NOT NULL;unique"  json:"username,omitempty"`                     //用户名
	Nickname        string     `gorm:"column:nickname" json:"nickname,omitempty"`                                      //昵称
	Password        string     `gorm:"column:password;NOT NULL" json:"password,omitempty"`                             //密码
	AvatarUrl       string     `gorm:"column:avatarUrl" json:"avatarUrl,omitempty"`                                    //头像URL
	Slogan          string     `gorm:"column:slogan" json:"slogan,omitempty"`                                          //标语
	CoverUrl        string     `gorm:"column:coverUrl" json:"coverUrl,omitempty"`                                      //封面URL
	CreatedAt       *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"` //创建时间
	UpdatedAt       *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`                           //更新时间
	EnableS3        string     `gorm:"column:enableS3;default:false;NOT NULL" json:"enableS3,omitempty"`
	Domain          string     `gorm:"column:domain" json:"domain,omitempty"`
	Bucket          string     `gorm:"column:bucket" json:"bucket,omitempty"`
	Region          string     `gorm:"column:region" json:"region,omitempty"`
	AccessKey       string     `gorm:"column:accessKey" json:"accessKey,omitempty"`
	SecretKey       string     `gorm:"column:secretKey" json:"secretKey,omitempty"`
	Endpoint        string     `gorm:"column:endpoint" json:"endpoint,omitempty"`
	ThumbnailSuffix string     `gorm:"column:thumbnailSuffix" json:"thumbnailSuffix,omitempty"`
	Favicon         string     `gorm:"column:favicon" json:"favicon,omitempty"`
	Title           string     `gorm:"column:title;default:'极简朋友圈';NOT NULL" json:"title,omitempty"`
	BeianNo         string     `gorm:"column:beianNo" json:"beianNo,omitempty"`
	Css             string     `gorm:"column:css" json:"css,omitempty"`
	Js              string     `gorm:"column:js" json:"js,omitempty"`
	Memos           []Memo     `json:"memos,omitempty"`
}

func (u *User) TableName() string {
	return "User"
}
