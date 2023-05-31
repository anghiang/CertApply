package utils

import (
	"context"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
)

func SendSingedTx(client *client.Client, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := client.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
