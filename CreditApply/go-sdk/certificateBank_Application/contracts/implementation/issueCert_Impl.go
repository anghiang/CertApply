package implementation

import (
	"context"
	"errors"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/contracts/interfaces"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var contractAddress = common.HexToAddress("0x904ed579402ed8bbb80ee7f0eb02e8226d78a70f")

// 合约实例 instance
var instance *interfaces.IssueCert

// 合约交易体
var session *interfaces.IssueCertSession

var clients *client.Client

func InitContract(_client *client.Client) {
	tempIns, err := interfaces.NewIssueCert(contractAddress, _client)
	if err != nil {
		fmt.Errorf("contract InitContract error %v", err)
	}
	tempSession := &interfaces.IssueCertSession{Contract: tempIns, CallOpts: *_client.GetCallOpts(), TransactOpts: *_client.GetTransactOpts()}
	instance = tempIns
	session = tempSession
	clients = _client
}
func AddRole(uAddr common.Address) bool {
	_, _, err := session.AddRole(uAddr)
	if err != nil {
		fmt.Errorf("contract AddRole error %v", err)
		return false
	}
	return true
}
func Issue(_hash [32]byte) (*types.Receipt, bool, error) {
	_, receipt, err := session.Issue(_hash)
	if err != nil {
		return nil, false, fmt.Errorf("contract issued error %v", err)

	} else if receipt.Output != "0x" {
		return nil, false, errors.New("该证书已存在!")
	}
	return receipt, true, nil
}

func Revoke(_hash [32]byte) (*types.Receipt, error, int) {
	_, receipt, err := session.Revoke(_hash)
	if err != nil {
		return nil, err, 400
	}
	if !IsIssued(_hash) {
		return nil, errors.New("该证书未被发布"), 401
	} else if receipt.Output != "0x0000000000000000000000000000000000000000000000000000000000000001" {
		return nil, errors.New("该证书已被撤销"), 402
	}
	return nil, nil, 200
	//if err != nil {
	//	fmt.Errorf("contract Revoke error %v", err)
	//	return nil, false, err
	//} else if receipt.Output != "0x0000000000000000000000000000000000000000000000000000000000000001" {
	//	return nil, false, errors.New("撤销失败")
	//}
	//return receipt, true, nil

}

func IsRole(_address common.Address) (bool, error) {
	address := "0x83309d045a19c44dc3722d15a6abd472f95866ac"
	judge, err := session.IsRole(common.HexToAddress(address))
	if err != nil {
		return false, err
	}
	return judge, nil
}

func IsIssued(_hash [32]byte) bool {
	judge, err := session.IsIssued(_hash)
	if err != nil || !judge {
		fmt.Errorf("contract IsIssued error %v", err)
		return false
	}
	return judge
}
func IsRevoked(_hash [32]byte) bool {
	judge, err := session.IsRevoked(_hash)
	if err != nil || !judge {
		fmt.Errorf("contract IsRevoked error %v", err)
		return false
	}
	return judge
}
func RevokeToBlkNum(_hash [32]byte) (*big.Int, error) {
	blockNum, err := session.CertRevoked(_hash)
	if err != nil {
		return nil, err
	}
	return blockNum, nil
}
func GetBlockByNumber(_blockNumber int64) (*types.Block, error) {
	blockInfo, err := clients.GetBlockByNumber(context.Background(), _blockNumber, true)
	if err != nil {
		return nil, err
	}
	return blockInfo, nil
}
