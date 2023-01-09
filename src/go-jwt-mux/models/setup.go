package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:gbqjWkd9!@tcp(localhost:3306)/go_jwt_mux"))
	if err != nil {
		fmt.Println("database connection failed")
	}

	db.AutoMigrate(&User{})

	DB = db
}
