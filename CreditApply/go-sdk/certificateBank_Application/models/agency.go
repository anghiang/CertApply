package models

import "github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"

// Agency 颁发机构
type Agency struct {
	Id         int
	Address    string
	AgencyName string
	password   string
}

func (a *Agency) QueryAddressByName() (string, error) {
	row := config.Db.QueryRow("select address from agency where mechanismName = ?", a.AgencyName)
	err := row.Scan(&a.Address)
	if err != nil {
		return "", err
	}
	defer config.Db.Close()
	return a.Address, nil
}
