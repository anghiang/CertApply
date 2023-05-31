package implementation

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

const (
	ISSUE     = "issue"
	REVOKE    = "revoke"
	ISROLE    = "isRole"
	ISISSUED  = "isIssued"
	ISREVOKED = "isRevoked"
)
const (
	DEFAULT_ACCOUNT = "0x83309d045a19c44dc3722d15a6abd472f95866ac"
)

type Contract struct {
	ContractAddress common.Address
	ContractABI     abi.ABI
}
type TransOtp struct {
	ChainId    *big.Int
	Nonce      *big.Int
	GasPrice   *big.Int
	GasLimit   *big.Int
	BlockLimit *big.Int
	GroupId    *big.Int
}

func (c *Contract) InitContract() {
	c.ContractAddress = common.HexToAddress("0x904ed579402ed8bbb80ee7f0eb02e8226d78a70f")
	abis := "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"document\",\"type\":\"bytes32\"}],\"name\":\"certIssuedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"document\",\"type\":\"bytes32\"}],\"name\":\"certRevokedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_role\",\"type\":\"address\"}],\"name\":\"addRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"certIssued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"certRevoked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cert\",\"type\":\"bytes32\"}],\"name\":\"isIssued\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cert\",\"type\":\"bytes32\"}],\"name\":\"isRevoked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cert\",\"type\":\"bytes32\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cert\",\"type\":\"bytes32\"}],\"name\":\"revoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	c.ContractABI, _ = abi.JSON(strings.NewReader(abis))
}
func (t *TransOtp) InitTransOpt(client *client.Client) {
	t.ChainId, _ = client.GetChainID(context.Background())
	t.GroupId = client.GetGroupID()
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(250), nil).Sub(max, big.NewInt(1))
	t.Nonce, _ = rand.Int(rand.Reader, max)
	t.GasPrice = big.NewInt(30000000)
	t.GasLimit = big.NewInt(30000000)
	blockNumber, _ := client.GetBlockNumber(context.Background())
	t.BlockLimit = big.NewInt(blockNumber + 600)
}
func SendSignedTx(signedTx *types.Transaction, client *client.Client) (*types.Receipt, error) {
	receipt, err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
func (c *Contract) Issue(_hash [32]byte, t *TransOtp) *types.Transaction {
	pack, err := c.ContractABI.Pack(ISSUE, _hash)
	if err != nil {
		fmt.Println("ContractABI.Pack ISSUE err :", err)
	}
	tx := types.NewTransaction(t.Nonce, c.ContractAddress, big.NewInt(0), t.GasLimit, t.GasPrice, t.BlockLimit, pack, t.ChainId, t.GroupId, nil, false)
	return tx
}
func (c *Contract) Revoke(_hash [32]byte, t *TransOtp) *types.Transaction {
	pack, err := c.ContractABI.Pack(REVOKE, _hash)
	if err != nil {
		fmt.Println("ContractABI.Pack REVOKE err :", err)
	}
	tx := types.NewTransaction(t.Nonce, c.ContractAddress, big.NewInt(0), t.GasLimit, t.GasPrice, t.BlockLimit, pack, t.ChainId, t.GroupId, nil, false)
	return tx
}
func (c *Contract) IsIssued(_hash [32]byte, clients *client.Client) (bool, error) {
	pack, err := c.ContractABI.Pack(ISISSUED, _hash)
	if err != nil {
		return false, err
	}
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress(DEFAULT_ACCOUNT),
		To:   &c.ContractAddress,
		Data: pack,
	}
	result, err := clients.CallContract(context.Background(), callMsg, nil)

	if err != nil {
		return false, err
	}
	return byteArrayToBool(result), nil
}

func (c *Contract) IsRevoked(_hash [32]byte, clients *client.Client) (bool, error) {
	pack, err := c.ContractABI.Pack(ISREVOKED, _hash)
	if err != nil {
		return false, err
	}
	callMsg := ethereum.CallMsg{
		To:   &c.ContractAddress,
		Data: pack,
	}
	result, err := clients.CallContract(context.Background(), callMsg, nil)

	if err != nil {
		return false, err
	}
	return byteArrayToBool(result), nil
}

func byteArrayToBool(byteArray []byte) bool {
	// 自定义转换规则，例如：只有当字节数组的最后一个字节为1时，将其转换为true
	if len(byteArray) > 0 && byteArray[len(byteArray)-1] == 1 {
		return true
	}
	return false
}
