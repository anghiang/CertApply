package transact

import (
	"WalletClient/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignTx(privateKey string, tx *types.Transaction) (*types.Transaction, error) {
	private, err := crypto.ToECDSA(common.Hex2Bytes(privateKey))
	//fmt.Println(private)
	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, private)
	if err != nil {
		fmt.Println("SignTx err ", err)
		return nil, err
	}
	return signTx, nil
}
