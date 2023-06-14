package models

import (
	"b-room/helper"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db_env := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(db_env))
	helper.PanicError(err)

	db.AutoMigrate(&User{}, &Category{}, &Room{}, &Reservation{})

	DB = db
}
