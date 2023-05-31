package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

const priKeyFile = "agencyKey"

func FindKeyFile(name string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	keyDir := filepath.Join(dir, priKeyFile)

	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		log.Fatal(err)
	}
	var keyFile string
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".pem") && strings.HasPrefix(f.Name(), name) {
			keyFile = filepath.Join(keyDir, f.Name())
			break
		}
	}

	if keyFile == "" {
		return "", fmt.Errorf("key file not found: %s", name)
	}
	return keyFile, nil
}
