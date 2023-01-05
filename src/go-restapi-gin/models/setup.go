package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//dsn := "host=localhost user=root password=gbqjWkd9! dbname=go_restapi_gin port=3306 sslmode=disable TimeZone=Asia/Seoul"
	database, err := gorm.Open(mysql.Open("root:gbqjWkd9!@tcp(localhost:3306)/go_restapi_gin"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
