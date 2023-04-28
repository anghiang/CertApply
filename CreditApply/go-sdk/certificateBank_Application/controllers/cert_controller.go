package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/services"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func VerifyCert(c *gin.Context) {
	//接收上传的文件
	file, _ := c.FormFile("cert")
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer f.Close()
	_certByte, _err := io.ReadAll(f)
	if _err != nil {
		c.JSON(http.StatusBadRequest, _err)
		return
	}
	certObj, verErr := services.VerifyCert(_certByte)
	if verErr != nil {
		c.JSON(http.StatusBadRequest, verErr.Error())
		return
	}
	fmt.Println("certObj : ", certObj)
	c.JSON(http.StatusOK, "ok")
}
func GetCertBlockNum(c *gin.Context) {
	var _tgtHashStr = c.PostForm("targetHash")
	_tgtHashStr = _tgtHashStr[2:]
	b, _ := hex.DecodeString(_tgtHashStr)
	var _tgtHashByt [32]byte
	copy(_tgtHashByt[32-len(b):], b)
	_, block, err := services.CertToBlockNum(_tgtHashByt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, block)
}

func Download(c *gin.Context) {
	var trans entity.TranscriptCert
	certName := c.PostForm("CertName")
	userName := c.PostForm("UserName")
	issueDate := c.PostForm("IssueDate")
	validityPeriod := c.PostForm("ValidityPeriod")
	agencyName := c.PostForm("AgencyName")
	signature := c.PostForm("Signature")
	certType := c.PostForm("CertType")
	CertNum := c.PostForm("CertNum")
	AgencyAddress := c.PostForm("AgencyAddress")
	hiddenData := c.PostFormArray("HiddenData")
	courses := c.PostFormArray("Course")

	trans.Transcript.Metadata.Number = CertNum
	trans.Transcript.Metadata.IssueDate = issueDate
	vp, err := strconv.Atoi(validityPeriod)
	if err != nil {
		fmt.Println(vp)
	}
	trans.Transcript.Metadata.ValidityPeriod = vp
	trans.Transcript.Metadata.CertName = certName
	trans.Transcript.Metadata.Signature = signature
	trans.Transcript.Agencies.AgencyName = agencyName
	trans.Transcript.Agencies.Address = AgencyAddress
	trans.Transcript.Users.UserName = userName

	//ua, err := c.Cookie("userAddr")
	//if err != nil {
	//	fmt.Println(err)
	//}
	ua := "iowaoiejfiwoejoifwoie"

	trans.Transcript.Users.Address = ua

	certTgtHash := trans.TranscriptTargetHash()
	trans.Transcript.Metadata.TargetHash = hex.EncodeToString(certTgtHash[:])
	ct, err := strconv.Atoi(certType)
	if err != nil {
		fmt.Println(vp)
	}
	trans.Transcript.CertType = ct

	var Courses []models.Course
	err = json.Unmarshal([]byte(courses[0]), &Courses)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return
	}
	trans.Transcript.Course = Courses

	var arr []string
	err = json.Unmarshal([]byte(hiddenData[0]), &arr)
	if err != nil {
		fmt.Println(err)
	}
	trans.HiddenData = arr

	//var _cert entity.Certs
	//
	certJson, err := trans.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}
	//生成.blkcert证书，并返回至前端
	certBytes, fileName := utils.NewFile(certJson)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("fileName", fileName)
	c.Writer.Write(certBytes)

	//c.JSON(http.StatusOK, gin.H{"download_status": "success"})
}
