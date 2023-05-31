package models

import (
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
)

// Agency 颁发机构
type Agency struct {
	Id         int    `json:"id"`
	Address    string `json:"address"`
	AgencyName string `json:"agency_name"`
	password   string `json:"password"`
}

func (a *Agency) QueryAddressByName() (int, string, error) {
	row := config.Db.QueryRow("select id,address from agency where mechanismName = ?", a.AgencyName)
	err := row.Scan(&a.Id, &a.Address)
	if err != nil {
		return 0, "", err
	}
	return a.Id, a.Address, nil
}

func (a *Agency) Login(pwd string) (string, error) {
	row := config.Db.QueryRow("select mechanismName from agency where address = ? and password = ?", a.Address, pwd)
	err := row.Scan(&a.AgencyName)
	if err != nil {
		return "", err
	}
	return a.AgencyName, nil
}
