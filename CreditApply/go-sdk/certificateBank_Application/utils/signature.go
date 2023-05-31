package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
)

func SignMsg(msg [32]byte, _privateKey *ecdsa.PrivateKey) ([]byte, error) {
	signByte, err := ecdsa.SignASN1(rand.Reader, _privateKey, msg[:])
	if err != nil {
		return nil, fmt.Errorf("sign message error : %v", err)
	}

	return signByte, nil
}
func VerifyMsg(msgHash [32]byte, signByte []byte, _publicKey *ecdsa.PublicKey) bool {
	if ecdsa.VerifyASN1(_publicKey, msgHash[:], signByte) {
		return true
	}
	return false
}
