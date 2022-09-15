package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

//连接数据库，创建表
func NewDB(filePath string) {
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		panic("sqlite 连接失败")
	}
	DB = db
}