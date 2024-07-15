package db

import (
	"github.com/kingwrcy/moments/vo"
	"github.com/samber/do/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(injector do.Injector) (*gorm.DB, error) {
	cfg := do.MustInvoke[vo.SysConfig](injector)
	db, err := gorm.Open(sqlite.Open(cfg.DB), &gorm.Config{})

	// 迁移 schema
	err = db.AutoMigrate(&User{}, &Comment{}, &Memo{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
