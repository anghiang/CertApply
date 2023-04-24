package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/services"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddTranscript(c *gin.Context) {
	validityPeriod := c.PostForm("ValidityPeriod")
	vpInt, err := strconv.Atoi(validityPeriod)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	transcript := entity.TranscriptCert{
		Transcript: models.Transcript{
			Metadata: models.Metadata{
				Id:             0,
				Number:         GenerateRandomString(8),
				IssueDate:      utils.GetTime(),
				ValidityPeriod: vpInt,
				CertName:       c.PostForm("CertName"),
			},
			Users: models.User{
				Address: c.PostForm("Address"),
			},
			CertType: 1,
		},
	}
	courses := c.PostFormArray("Courses")
	var Courses []models.Course
	err = json.Unmarshal([]byte(courses[0]), &Courses)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"add_status": "add_err",
		})
		fmt.Print("error : ", err)
		return
	}
	transcript.Transcript.Course = Courses

	transcript.Transcript.Agencies.AgencyName = "清华大学"
	//transcript.Transcript.Agencies.Address, err = transcript.Transcript.Agencies.QueryAddressByName()
	transcript.Transcript.Agencies.Address = "0x83309d045a19c44dc3722d15a6abd472f95866ac"
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"add_status": "add_err",
		})
		fmt.Print("error : ", err)
		return
	}
	transcript.Transcript.Users.UserName, err = transcript.Transcript.Users.QueryUNameByAddress()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"add_status": "add_err",
		})
		fmt.Print("error : ", err)
		return
	}
	certTgtHash := transcript.TranscriptTargetHash()
	err = services.IssueTranscript(transcript, certTgtHash)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"add_status": "add_err",
		})
		fmt.Print("error : ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"add_status": "add_ok",
	})
	defer config.Db.Close()
}
