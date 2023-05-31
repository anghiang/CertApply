package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/contracts/implementation"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/services"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func VerifySig(c *gin.Context) {
	targetHash := c.PostForm("targetHash")
	signature := c.PostForm("signature")
	address := c.PostForm("address")
	result, err := services.VerySignature(targetHash, signature, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if result {
		c.JSON(200, "验证成功")
		return
	}
	c.JSON(http.StatusBadRequest, "验证失败")

}

func VerifyCert(c *gin.Context) {
	var cert *entity.Certs
	err := c.BindJSON(&cert)
	if err != nil {
		fmt.Println(err)
	}
	err = utils.BindOpts("0x83309d045a19c44dc3722d15a6abd472f95866ac")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	_, verErr := services.VerifyCert(cert)

	tempHsh, _ := hex.DecodeString(cert.Cert.Metadata.TargetHash)
	var hashByte [32]byte
	copy(hashByte[:], tempHsh)

	_, block, err := services.RevokeToBlockNum(hashByte)
	if err != nil {
		fmt.Println(err)
	}

	if verErr != nil {
		fmt.Println("verErr", verErr)
		c.JSON(200, gin.H{
			"verify_result": "verify_err",
			"verify_status": implementation.IsRevoked(hashByte),
			"block":         block,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("certificate", cert)
	err = session.Save()
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, gin.H{
		"verify_result": "verify_ok",
	})
}

func VerifyTrans(c *gin.Context) {
	var trans *entity.TranscriptCert
	err := c.BindJSON(&trans)
	if err != nil {
		fmt.Println(err)
	}
	err = utils.BindOpts("0x83309d045a19c44dc3722d15a6abd472f95866ac")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	_, verErr := services.VerifyTrans(trans)

	tempHsh, _ := hex.DecodeString(trans.Transcript.Metadata.TargetHash)
	var hashByte [32]byte
	copy(hashByte[:], tempHsh)

	_, block, err := services.RevokeToBlockNum(hashByte)

	if block != nil {
		timestampHex, _ := strconv.ParseInt(block.Timestamp[2:], 16, 64)
		//fmt.Println(timestampHex)
		fmt.Println(block.Timestamp[2:10])
		t := time.Unix(timestampHex/1000, 0)
		block.Timestamp = t.Format("2006-01-02 15:04:05")
	}
	if err != nil {
		fmt.Println(err)
	}

	if verErr != nil {

		c.JSON(200, gin.H{
			"verify_result": "verify_err",
			"verify_status": implementation.IsRevoked(hashByte),
			"block":         block,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("certificate", trans)
	err = session.Save()
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, gin.H{
		"verify_result": "verify_ok",
	})
}

func RevokeCert(c *gin.Context) {
	targetHash := c.PostForm("revokeHash")
	session := sessions.Default(c)
	agencyAddr := session.Get(AgencyAddr)
	//通过已登录的账号绑定TransOpts和CallOpts
	err := utils.BindOpts(agencyAddr.(string))
	if err != nil {
		fmt.Println("bind err : ", err)
		c.JSON(200, err.Error())
		return
	}
	targetHash = targetHash[2:]
	hashByte, err := hex.DecodeString(targetHash)
	if err != nil {
		fmt.Println("hex err : ", err)
		c.JSON(200, gin.H{
			"revoke_status": "revoke_err",
		})
		return
	}
	var _tgtHashByt [32]byte
	copy(_tgtHashByt[32-len(hashByte):], hashByte)
	err, code := services.RevokeCert(_tgtHashByt)
	fmt.Println("code: ", code)
	if code == 400 {
		fmt.Println("RevokeCert err")
	}
	c.JSON(200, gin.H{
		"revoke_status": code,
	})
	return
}

//func GetCertBlockNum(c *gin.Context) {
//	var _tgtHashStr = c.PostForm("targetHash")
//	_tgtHashStr = _tgtHashStr[2:]
//	b, _ := hex.DecodeString(_tgtHashStr)
//	var _tgtHashByt [32]byte
//	copy(_tgtHashByt[32-len(b):], b)
//	_, block, err := services.CertToBlockNum(_tgtHashByt)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	c.JSON(200, block)
//}

func DownloadTrans(c *gin.Context) {
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

	session := sessions.Default(c)
	userAddr := session.Get(UserAddr)

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

	trans.HiddenData = entity.CalHidHash(arr)
	fmt.Println(trans.HiddenData)

	trans.Transcript.Users.Address = userAddr.(string)
	certTgtHash := trans.TranscriptTargetHash()
	trans.Transcript.Metadata.TargetHash = hex.EncodeToString(certTgtHash[:])

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
}

func DownloadCert(c *gin.Context) {
	var cert entity.Certs

	certName := c.PostForm("CertName")
	userName := c.PostForm("UserName")
	issueDate := c.PostForm("IssueDate")
	validityPeriod := c.PostForm("ValidityPeriod")
	agencyName := c.PostForm("AgencyName")
	signature := c.PostForm("Signature")
	certType := c.PostForm("CertType")
	CertNum := c.PostForm("CertNum")
	AgencyAddress := c.PostForm("AgencyAddress")

	cert.Cert.Metadata.Number = CertNum
	cert.Cert.Metadata.IssueDate = issueDate
	vp, err := strconv.Atoi(validityPeriod)
	if err != nil {
		fmt.Println(vp)
	}
	cert.Cert.Metadata.ValidityPeriod = vp
	cert.Cert.Metadata.CertName = certName
	cert.Cert.Metadata.Signature = signature
	cert.Cert.Agencies.AgencyName = agencyName
	cert.Cert.Agencies.Address = AgencyAddress
	cert.Cert.Users.UserName = userName

	session := sessions.Default(c)
	userAddr := session.Get(UserAddr)
	if err != nil {
		fmt.Println(err)
	}

	ct, err := strconv.Atoi(certType)
	if err != nil {
		fmt.Println(vp)
	}
	cert.Cert.CertType = ct

	cert.Cert.Users.Address = userAddr.(string)

	certTgtHash := cert.TargetHash()
	cert.Cert.Metadata.TargetHash = hex.EncodeToString(certTgtHash[:])

	certJson, err := cert.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}
	//生成.blkcert证书，并返回至前端
	certBytes, fileName := utils.NewFile(certJson)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("fileName", fileName)
	c.Writer.Write(certBytes)
}
