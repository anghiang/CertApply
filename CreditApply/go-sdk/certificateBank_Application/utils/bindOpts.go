package utils

import "github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"

func BindOpts(_address string) error {
	keyFile, err := FindKeyFile(_address)
	if err != nil {
		return err
	}
	err = config.InitClient(keyFile)
	if err != nil {
		return err
	}
	return nil

}
