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
		if fieldName == "Signature" {
			continue
		}
		fieldValue := MetadataV.Field(i).Interface()
		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", fieldName, fieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}

	AgenciesV := reflect.ValueOf(t.Transcript.Agencies)
	AgenciesT := AgenciesV.Type()
	for i := 0; i < AgenciesT.NumField(); i++ {
		AFieldName := AgenciesT.Field(i).Name
		if AFieldName == "password" || AFieldName == "Id" {
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
		if UFieldName == "password" || UFieldName == "Id" {
			continue
		}
		UFieldValue := UserV.Field(i).Interface()

		dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", UFieldName, UFieldValue, salt)), "\n")))
		unSortedHash = append(unSortedHash, dataHash)
	}

	for _, tc := range t.Transcript.Course {
		CourseV := reflect.ValueOf(tc)
		CourseT := CourseV.Type()
		for i := 0; i < CourseT.NumField(); i++ {
			CFieldName := CourseT.Field(i).Name
			if CFieldName == "password" || CFieldName == "Id" {
				continue
			}
			CFieldValue := CourseV.Field(i).Interface()

			dataHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", CFieldName, CFieldValue, salt)), "\n")))
			unSortedHash = append(unSortedHash, dataHash)
		}
	}
	ctHash := sha256.Sum256([]byte(strings.TrimSuffix(strings.TrimSpace(fmt.Sprintf("%s:%v%v", "CertType", t.Transcript.CertType, salt)), "\n")))
	unSortedHash = append(unSortedHash, ctHash)

	if len(t.HiddenData) != 0 {
		for i := 0; i < len(t.HiddenData); i++ {
			var arr [32]byte
			copy(arr[:], t.HiddenData[i])
			unSortedHash = append(unSortedHash, arr)
		}
	}
	return unSortedHash
}

func (t *TranscriptCert) TranscriptTargetHash() [32]byte {
	hash := t.CalTranscriptFieldHash()
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

func (t *TranscriptCert) MarshalJSON() ([]byte, error) {
	var tmpCourses []interface{}
	for _, c := range t.Transcript.Course {
		fmt.Println(c.Score)
		fmt.Println(c.CourseName)
		if c.CourseName == "" && c.Score != 0 {
			tmpCourses = append(tmpCourses, struct {
				Score float64
			}{c.Score})
			continue
		} else if c.CourseName != "" && c.Score == 0 {
			tmpCourses = append(tmpCourses, struct {
				CourseName string
			}{c.CourseName})
			continue
		} else {
			tmpCourses = append(tmpCourses, struct {
				CourseName string
				Score      float64
			}{c.CourseName,
				c.Score,
			})
		}
	}

	return json.MarshalIndent(struct {
		Transcript struct {
			Metadata struct {
				Number         string `json:"number"`
				IssueDate      string `json:"issue_date"`
				ValidityPeriod int    `json:"validity_period"`
				Signature      string `json:"signature"`
				CertName       string `json:"cert_name"`
				TargetHash     string `json:"target_hash"`
			} `json:"metadata"`

			Course []interface{} `json:"course"`

			Users struct {
				UserName string `json:"user_name"`
				Address  string `json:"address"`
			} `json:"users"`

			Agencies struct {
				Address    string `json:"address"`
				AgencyName string `json:"agency_name"`
			} `json:"agencies"`

			CertType   int      `json:"cert_type"`
			HiddenData []string `json:"hidden_data"`
		}
	}{
		Transcript: struct {
			Metadata struct {
				Number         string `json:"number"`
				IssueDate      string `json:"issue_date"`
				ValidityPeriod int    `json:"validity_period"`
				Signature      string `json:"signature"`
				CertName       string `json:"cert_name"`
				TargetHash     string `json:"target_hash"`
			} `json:"metadata"`

			Course []interface{} `json:"course"`

			Users struct {
				UserName string `json:"user_name"`
				Address  string `json:"address"`
			} `json:"users"`

			Agencies struct {
				Address    string `json:"address"`
				AgencyName string `json:"agency_name"`
			} `json:"agencies"`

			CertType   int      `json:"cert_type"`
			HiddenData []string `json:"hidden_data"`
		}{
			Metadata: struct {
				Number         string `json:"number"`
				IssueDate      string `json:"issue_date"`
				ValidityPeriod int    `json:"validity_period"`
				Signature      string `json:"signature"`
				CertName       string `json:"cert_name"`
				TargetHash     string `json:"target_hash"`
			}{
				Number:         t.Transcript.Metadata.Number,
				IssueDate:      t.Transcript.Metadata.IssueDate,
				ValidityPeriod: t.Transcript.Metadata.ValidityPeriod,
				Signature:      t.Transcript.Metadata.Signature,
				CertName:       t.Transcript.Metadata.CertName,
				TargetHash:     t.Transcript.Metadata.TargetHash,
			},
			Course: tmpCourses,
			Users: struct {
				UserName string `json:"user_name"`
				Address  string `json:"address"`
			}{
				UserName: t.Transcript.Users.UserName,
				Address:  t.Transcript.Users.Address,
			},
			Agencies: struct {
				Address    string `json:"address"`
				AgencyName string `json:"agency_name"`
			}{
				Address:    t.Transcript.Agencies.Address,
				AgencyName: t.Transcript.Agencies.AgencyName,
			},
			CertType:   t.Transcript.CertType,
			HiddenData: t.HiddenData,
		}},
		"", " ",
	)
}
