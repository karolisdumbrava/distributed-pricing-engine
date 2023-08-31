package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPass, DbHost, DbPort, DbName)
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}