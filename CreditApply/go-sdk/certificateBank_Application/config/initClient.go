package config

import (
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/contracts/implementation"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

func InitClient(keyFile string) error {

	configs, err := conf.ParseConfigFile("config.toml", keyFile)

	if err != nil {
		return err
	}
	config := &configs[0]
	clients, err := client.Dial(config)
	if err != nil {
		return err
	}
	implementation.InitContract(clients)
	return nil
}
