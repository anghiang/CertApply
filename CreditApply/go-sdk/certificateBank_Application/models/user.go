package models

import (
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
)

// User 用户
type User struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Address  string `json:"address"`
	password string `json:"password"`
}

func QueryUidByAddress(address string) (uid int64, err error) {
	row := config.Db.QueryRow("select id from users where address = ?", address)
	err = row.Scan(&uid)
	if err != nil {
		return -1, err
	}
	return uid, nil
}

func (u *User) QueryUNameByAddress() (name string, err error) {
	row := config.Db.QueryRow("select name from users where address =?", u.Address)
	err = row.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (u *User) Login(pwd string) (name string, err error) {
	row := config.Db.QueryRow("select name from users where address = ? and password = ?", u.Address, pwd)
	err = row.Scan(&u.UserName)
	if err != nil {
		return "", err
	}
	return u.UserName, nil
}
