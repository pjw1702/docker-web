package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB open
func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "gbqjWkd9!"
	dbName := "go_crud"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
