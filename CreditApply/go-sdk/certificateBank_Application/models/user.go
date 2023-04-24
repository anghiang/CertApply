package models

import "github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"

// User 用户
type User struct {
	Id       int
	UserName string
	Address  string
	password string
}

func QueryUidByAddress(address string) (uid int64, err error) {
	row := config.Db.QueryRow("select id from users where address = ?", address)
	err = row.Scan(&uid)
	if err != nil {
		panic(err)
		return -1, err
	}
	defer config.Db.Close()
	return uid, nil
}

func QueryUNameByAddress(address string) (name string, err error) {
	row := config.Db.QueryRow("select name from users where address = ?", address)
	err = row.Scan(&name)
	if err != nil {
		return "", err
	}
	defer config.Db.Close()
	return name, nil
}
