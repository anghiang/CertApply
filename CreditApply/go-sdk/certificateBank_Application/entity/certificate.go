package entity

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"reflect"
	"sort"
	"strings"
)

type Certs struct {
	Cert       models.Cert
	HiddenData []string
}

func NewCerts(cert models.Cert) *Certs {
	return &Certs{Cert: cert}
}

func (c *Certs) UnMarshalJSON(data []byte) error {
	err := json.Unmarshal(data, c)
	if err != nil {
		return fmt.Errorf("Error Unmarshalling JSON: %v", err)
	}
	return nil
}
func (c *Certs) MarshalJSON() ([]byte, error) {
	data, err := json.MarshalIndent(*c, "", " ")
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)
	}
	return data, nil
}
func (c *Certs) CalFieldHash() (unSortedHash [][32]byte) {
	const salt = "jhblockchain"
	MetadataV := reflect.ValueOf(c.Cert.Metadata)
	MetadataT := MetadataV.Type()
	for i := 0; i < MetadataT.NumField(); i++ {
		fieldName := MetadataT.Field(i).Name
		fieldValue := MetadataV.Field(i).Interface()
		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", fieldName, fieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}
	AgenciesV := reflect.ValueOf(c.Cert.Agencies)
	AgenciesT := AgenciesV.Type()
	for i := 0; i < AgenciesT.NumField(); i++ {
		AFieldName := AgenciesT.Field(i).Name
		if AFieldName == "password" {
			continue
		}
		AFieldValue := AgenciesV.Field(i).Interface()

		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", AFieldName, AFieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}

	UserV := reflect.ValueOf(c.Cert.Users)
	UserT := AgenciesV.Type()
	for i := 0; i < UserT.NumField(); i++ {
		UFieldName := UserT.Field(i).Name
		if UFieldName == "password" {
			continue
		}
		UFieldValue := UserV.Field(i).Interface()

		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", UFieldName, UFieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}

	signHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", "CertType", c.Cert.CertType, salt)), "\n")))
	unSortedHash = append(unSortedHash, signHash)

	if len(c.HiddenData) != 0 {
		for i := 0; i < len(c.HiddenData); i++ {
			var arr [32]byte
			copy(arr[:], c.HiddenData[i])
			unSortedHash = append(unSortedHash, arr)
		}
	}
	return unSortedHash
}

func (c *Certs) TargetHash() [32]byte {
	hash := c.CalFieldHash()
	sort.Slice(hash, func(i, j int) bool {
		return bytes.Compare(hash[i][:], hash[j][:]) == -1
	})
	var temp string
	for _, v := range hash {
		temp += hex.EncodeToString(v[:])

	}
	b := sha256.Sum256([]byte(temp))
	return b
}
