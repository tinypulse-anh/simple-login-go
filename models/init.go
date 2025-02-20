package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Init() {
	if DB != nil {
		return
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err == nil {
		DB = db

		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(100)
	} else {
		panic(err)
	}
}
