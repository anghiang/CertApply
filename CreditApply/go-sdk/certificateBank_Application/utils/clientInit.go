package utils

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/contracts/implementation"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

func Init() (*client.Client, *implementation.Contract, *implementation.TransOtp) {
	configs, err := conf.ParseConfigFile("config.toml", ".ci/0x83309d045a19c44dc3722d15a6abd472f95866ac.pem")

	if err != nil {
		fmt.Println("ParseConfigFile err ", err)
		return nil, nil, nil
	}
	config := &configs[0]
	clients, err := client.Dial(config)
	if err != nil {
		fmt.Println("Dial err ", err)
		return nil, nil, nil
	}
	contract := new(implementation.Contract)
	contract.InitContract()
	transOpts := new(implementation.TransOtp)
	transOpts.InitTransOpt(clients)
	return clients, contract, transOpts
}
