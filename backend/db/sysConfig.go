package db

type SysConfig struct {
	Id      int32  `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	Content string `gorm:"column:content;NOT NULL;default:'{}'"  json:"content,omitempty"`
}

func (s *SysConfig) TableName() string {
	return "SysConfig"
}
