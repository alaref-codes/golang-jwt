package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	DBConn = db
	if err != nil {
		panic("failed to connect database")
	}
	DBConn.AutoMigrate(&User{})
}
