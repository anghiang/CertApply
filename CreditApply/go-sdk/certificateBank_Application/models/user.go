package models

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
)

// User 用户
type User struct {
	Id       int
	UserName string
	Address  string
	password string
}

func QueryUidByAddress(address string) (uid int64, err error) {
	fmt.Println(address)
	row := config.Db.QueryRow("select id from users where address = ?", address)
	err = row.Scan(&uid)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return uid, nil
}

func QueryUNameByAddress(address string) (name string, err error) {
	row := config.Db.QueryRow("select name from users where address = ?", address)
	err = row.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}
