package models

import (
	"database/sql"

	"github.com/pjw1702/go-auth/config"
	entites "github.com/pjw1702/go-auth/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entites.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("select id, full_name, email, username, password from users where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.Full_Name, &user.Email, &user.Username, &user.Password)
	}

	return nil

}
