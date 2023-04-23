package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"time"
)

const fileSuffix = ".blkcert"

var fileName = strconv.FormatInt(time.Now().Unix(), 10)

func NewFile(data []byte) ([]byte, string) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	w.Write(data)
	w.Flush()

	filenameWithExt := fmt.Sprintf("%s%s", fileName, fileSuffix)

	return buf.Bytes(), filenameWithExt
}
