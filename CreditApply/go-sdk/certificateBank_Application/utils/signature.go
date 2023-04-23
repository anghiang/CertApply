package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func SignMsg(msg []byte, _privateKey *ecdsa.PrivateKey) ([]byte, error) {
	messageHash := sha256.Sum256(msg)
	signByte, err := ecdsa.SignASN1(rand.Reader, _privateKey, messageHash[:])
	if err != nil {
		return nil, fmt.Errorf("sign message error : %v", err)
	}

	return signByte, nil
}
func VerifyMsg(msgHash, signByte []byte, _publicKey *ecdsa.PublicKey) bool {
	if ecdsa.VerifyASN1(_publicKey, msgHash[:], signByte) {
		return true
	}
	return false
}
