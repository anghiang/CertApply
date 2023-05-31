package entity

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"reflect"
	"sort"
	"strings"
)

func init() {
	gob.Register(&Certs{})
}

type Certs struct {
	Cert models.Cert
}

func (c *Certs) UnMarshalJSON(data []byte) error {
	err := json.Unmarshal(data, c)
	if err != nil {
		return fmt.Errorf("Error Unmarshalling JSON: %v", err)
	}
	return nil
}

func (c *Certs) CalFieldHash() (unSortedHash [][32]byte) {
	const salt = "jhblockchain"
	MetadataV := reflect.ValueOf(c.Cert.Metadata)
	MetadataT := MetadataV.Type()
	for i := 0; i < MetadataT.NumField(); i++ {
		fieldName := MetadataT.Field(i).Name
		if fieldName == "Signature" || fieldName == "TargetHash" || fieldName == "Id" {
			continue
		}
		fieldValue := MetadataV.Field(i).Interface()
		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", fieldName, fieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}
	AgenciesV := reflect.ValueOf(c.Cert.Agencies)
	AgenciesT := AgenciesV.Type()
	for i := 0; i < AgenciesT.NumField(); i++ {
		AFieldName := AgenciesT.Field(i).Name
		if AFieldName == "password" || AFieldName == "Id" {
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
		if UFieldName == "password" || UFieldName == "Id" {
			continue
		}
		UFieldValue := UserV.Field(i).Interface()

		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", UFieldName, UFieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}

	ctHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", "CertType", c.Cert.CertType, salt)), "\n")))
	unSortedHash = append(unSortedHash, ctHash)

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

func (c *Certs) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Cert struct {
			Metadata struct {
				Number         string `json:"number"`
				IssueDate      string `json:"issue_date"`
				ValidityPeriod int    `json:"validity_period"`
				Signature      string `json:"signature"`
				CertName       string `json:"cert_name"`
				TargetHash     string `json:"target_hash"`
			} `json:"metadata"`

			Users struct {
				UserName string `json:"user_name"`
				Address  string `json:"address"`
			} `json:"users"`

			Agencies struct {
				Address    string `json:"address"`
				AgencyName string `json:"agency_name"`
			} `json:"agencies"`

			CertType int `json:"cert_type"`
		}
	}{
		Cert: struct {
			Metadata struct {
				Number         string `json:"number"`
				IssueDate      string `json:"issue_date"`
				ValidityPeriod int    `json:"validity_period"`
				Signature      string `json:"signature"`
				CertName       string `json:"cert_name"`
				TargetHash     string `json:"target_hash"`
			} `json:"metadata"`

			Users struct {
				UserName string `json:"user_name"`
				Address  string `json:"address"`
			} `json:"users"`

			Agencies struct {
				Address    string `json:"address"`
				AgencyName string `json:"agency_name"`
			} `json:"agencies"`

			CertType int `json:"cert_type"`
		}{
			Metadata: struct {
				Number         string `json:"number"`
				IssueDate      string `json:"issue_date"`
				ValidityPeriod int    `json:"validity_period"`
				Signature      string `json:"signature"`
				CertName       string `json:"cert_name"`
				TargetHash     string `json:"target_hash"`
			}{
				Number:         c.Cert.Metadata.Number,
				IssueDate:      c.Cert.Metadata.IssueDate,
				ValidityPeriod: c.Cert.Metadata.ValidityPeriod,
				Signature:      c.Cert.Metadata.Signature,
				CertName:       c.Cert.Metadata.CertName,
				TargetHash:     c.Cert.Metadata.TargetHash,
			},
			Users: struct {
				UserName string `json:"user_name"`
				Address  string `json:"address"`
			}{
				UserName: c.Cert.Users.UserName,
				Address:  c.Cert.Users.Address,
			},
			Agencies: struct {
				Address    string `json:"address"`
				AgencyName string `json:"agency_name"`
			}{
				Address:    c.Cert.Agencies.Address,
				AgencyName: c.Cert.Agencies.AgencyName,
			},
			CertType: c.Cert.CertType,
		}},
		"", " ",
	)
}
