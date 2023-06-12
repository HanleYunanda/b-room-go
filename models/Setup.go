package models

import (
	"b-room/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/b-room"))
	helper.PanicError(err)

	db.AutoMigrate(&User{}, &Category{}, &Room{}, &Reservation{})

	DB = db
}
