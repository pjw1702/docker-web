package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:gbqjWkd9!@tcp(localhost:3306)/go_restapi_mux"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
	DB = db
}
