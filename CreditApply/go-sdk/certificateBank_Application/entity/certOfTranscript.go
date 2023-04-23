package entity

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"reflect"
	"sort"
	"strings"
)

type TranscriptCert struct {
	Transcript models.Transcript
	HiddenData []string
}

func (t *TranscriptCert) CalTranscriptFieldHash() (unSortedHash [][32]byte) {
	const salt = "jhblockchain"
	MetadataV := reflect.ValueOf(t.Transcript.Metadata)
	MetadataT := MetadataV.Type()
	for i := 0; i < MetadataT.NumField(); i++ {
		fieldName := MetadataT.Field(i).Name
		fieldValue := MetadataV.Field(i).Interface()
		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", fieldName, fieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}

	AgenciesV := reflect.ValueOf(t.Transcript.Agencies)
	AgenciesT := AgenciesV.Type()
	for i := 0; i < AgenciesT.NumField(); i++ {
		AFieldName := AgenciesT.Field(i).Name
		if AFieldName == "password" {
			continue
		}
		AFieldValue := AgenciesV.Field(i).Interface()

		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("\"%s\":\"%v\"%v", AFieldName, AFieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}

	UserV := reflect.ValueOf(t.Transcript.Users)
	UserT := AgenciesV.Type()
	for i := 0; i < UserT.NumField(); i++ {
		UFieldName := UserT.Field(i).Name
		if UFieldName == "password" {
			continue
		}
		UFieldValue := UserV.Field(i).Interface()

		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("\"%s\":\"%v\"%v", UFieldName, UFieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}

	//for _, tc := range t.Transcript.Course {
	//	CourseV := reflect.ValueOf(tc)
	//	CourseT := CourseV.Type()
	//	for i := 0; i < CourseT.NumField(); i++ {
	//		CFieldName := CourseT.Field(i).Name
	//		if CFieldName == "password" {
	//			continue
	//		}
	//		CFieldValue := CourseV.Field(i).Interface()
	//
	//		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("\"%s\":\"%v\"%v", UFieldName, UFieldValue, salt)), "\n")))
	//		unSortedHash = append(unSortedHash, dataHash)
	//	}
	//}

	if len(t.HiddenData) != 0 {
		for i := 0; i < len(t.HiddenData); i++ {
			var arr [32]byte
			copy(arr[:], t.HiddenData[i])
			unSortedHash = append(unSortedHash, arr)
		}
	}
	return unSortedHash
}

func (c *Certs) TranscriptTargetHash() [32]byte {
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
