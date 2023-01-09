package models

type User struct {
	Id       int64  `gorm:"primarykey" json:"id"`
	FullName string `gorm:"varchar(300)" json:"full_name"`
	Username string `gorm:"varchar(300)" json:"username"`
	Password string `gorm:"varchar(300)" json:"password"`
}
