package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/services"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
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
	certName := c.PostForm("CertName")
	userName := c.PostForm("UserName")
	issueDate := c.PostForm("IssueDate")
	validityPeriod := c.PostForm("ValidityPeriod")
	hiddenData := c.PostFormArray("HiddenData")
	courses := c.PostFormArray("Course")
	fmt.Println(certName)
	fmt.Println(userName)
	fmt.Println(issueDate)
	fmt.Println(validityPeriod)
	fmt.Println(courses[0])
	fmt.Println(hiddenData[0])

	var Courses []models.Course
	err := json.Unmarshal([]byte(courses[0]), &Courses)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return
	}

	for _, c := range Courses {
		fmt.Printf("CourseName: %s, Score: %.2f\n", c.CourseName, c.Score)
	}

	var arr []string
	err = json.Unmarshal([]byte(hiddenData[0]), &arr)
	if err != nil {
		fmt.Println(err)
	}
	//var _cert entity.Certs
	//
	//_, err := _cert.MarshalJSON()
	//if err != nil {
	//	return err
	//}
	////生成.blkcert证书，并返回至前端
	//certBytes, fileName := utils.NewFile(certJson)
	////将证书数据存储至mysql数据库
	//err = _cert.Cert.AddCert()
	//if err != nil {
	//}
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//c.Header("Content-Type", "application/octet-stream")
	//c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	//c.Writer.Write(certBytes)
	//c.JSON(http.StatusOK, gin.H{"message": "success"})
}
