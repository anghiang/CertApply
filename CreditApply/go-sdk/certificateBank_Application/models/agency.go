package models

import "github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"

// Agency 颁发机构
type Agency struct {
	Id         int
	Address    string
	AgencyName string
	password   string
}

func QueryIdByAddress(address string) (uid int64, err error) {
	row := config.Db.QueryRow("select id from agency where address = ?", address)
	err = row.Scan(&uid)
	if err != nil {
		panic(err)
		return -1, err
	}
	defer config.Db.Close()
	return uid, nil
}
