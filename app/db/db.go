package db

import (
	"gorm.io/gorm"
)

// db 接口
type IDb interface {
	Dsn() string
	Dialector() gorm.Dialector
	NewConnInfo() IDb
}

// 初始化資料庫連線
func InitDb(i IDb) *gorm.DB {
	dialector := i.Dialector()
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
