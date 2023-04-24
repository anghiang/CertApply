package utils

import (
	"bytes"
	"crypto/md5"
)

func ByteTrans() []byte {
	var hashChannel = make(chan []byte, 1)
	var buffer bytes.Buffer
	sum := md5.Sum(buffer.Bytes())
	hashChannel <- sum[:]
	return <-hashChannel
}
