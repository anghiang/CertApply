package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetPrivateKey(accountKeyFile string) (*ecdsa.PrivateKey, error) {
	privateKey, _, err := conf.LoadECPrivateKeyFromPEMS(accountKeyFile)
	if err != nil {
		return nil, fmt.Errorf("parse private key failed, err: %v", err)
	}
	return privateKey, nil
}
func PublicKeyToAddr(pubKey ecdsa.PublicKey) common.Address {
	pubKeyBytes := elliptic.Marshal(pubKey.Curve, pubKey.X, pubKey.Y)
	addressBytes := crypto.Keccak256(pubKeyBytes[1:])
	address := common.HexToAddress(hex.EncodeToString(addressBytes[12:]))
	return address
}
