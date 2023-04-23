package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Init() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/blockcertapplication"
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Print(err)
	}
}
