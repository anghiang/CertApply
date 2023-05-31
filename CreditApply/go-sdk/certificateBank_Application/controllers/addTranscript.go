package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/services"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddTranscriptSign(c *gin.Context) {
	validityPeriod := c.PostForm("ValidityPeriod")
	vpInt, err := strconv.Atoi(validityPeriod)
	if err != nil {
		//c.JSON(200, gin.H{
		//	"add_status": "add_err",
		//})
		fmt.Print("error : ", err)
	}
	transcript := entity.TranscriptCert{
		Transcript: models.Transcript{
			Metadata: models.Metadata{
				Number:         GenerateRandomString(8),
				IssueDate:      utils.GetTime(),
				ValidityPeriod: vpInt,
				CertName:       c.PostForm("CertName"),
			},
			Users: models.User{
				Address: c.PostForm("Address"),
			},
			CertType: 2,
		},
	}
	courses := c.PostFormArray("Courses")
	var Courses []models.Course
	err = json.Unmarshal([]byte(courses[0]), &Courses)
	if err != nil {
		//c.JSON(200, gin.H{
		//	"add_status": "add_err",
		//})
		fmt.Print("error : ", err)
		return
	}
	transcript.Transcript.Course = Courses
	session := sessions.Default(c)
	agencyName := session.Get(AgencyName)
	transcript.Transcript.Agencies.AgencyName = agencyName.(string)
	transcript.Transcript.Agencies.Id, transcript.Transcript.Agencies.Address, err = transcript.Transcript.Agencies.QueryAddressByName()
	if err != nil {
		//c.JSON(200, gin.H{
		//	"add_status": "add_err",
		//})
		fmt.Print("error : ", err)
		return
	}
	transcript.Transcript.Users.UserName, err = transcript.Transcript.Users.QueryUNameByAddress()
	if err != nil {
		//c.JSON(200, gin.H{
		//	"add_status": "add_err",
		//})
		fmt.Print("error : ", err)
		return
	}

	certTgtHash := transcript.TranscriptTargetHash()
	cli, contract, tr = utils.Init()
	tx := contract.Issue(certTgtHash, tr)
	signature, err := services.SignCertTgtHsh(certTgtHash, transcript.Transcript.Agencies.Address)
	transcript.Transcript.Metadata.Signature = signature
	transcript.Transcript.AddTranscript()

	rawTxByte, err := tx.MarshalJSON()
	fmt.Println("rawTxByte: ", rawTxByte)
	if err != nil {
		fmt.Println("tx.MarshalJson err ", err)
		return
	}
	var url = "http://127.0.0.1:5000/setSignature"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(rawTxByte))
	if err != nil {
		// 处理创建请求对象时的错误
		panic(err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer your_token")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// 处理请求发送失败的错误
		panic(err)
	}
	defer resp.Body.Close()
	//err = services.IssueTranscript(transcript, certTgtHash)
	//if err != nil {
	//	c.JSON(200, gin.H{
	//		"add_status": "add_err",
	//	})
	//	fmt.Print("error : ", err)
	//	return
	//}
	//c.JSON(200, gin.H{
	//	"add_status": "add_ok",
	//})
}
