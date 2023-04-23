package controllers

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func AddCertShow(c *gin.Context) {
	c.HTML(http.StatusOK, "issueCert.html", nil)
}

func AddCert(c *gin.Context) {
	var cert entity.Certs
	certName := c.PostForm("CertName")
	address := c.PostForm("Address")
	validityPeriod := c.PostForm("ValidityPeriod")

	cert.Cert.CertType = 1
	cert.Cert.Metadata.CertName = certName
	cert.Cert.Metadata.Number = generateRandomString(8)
	cert.Cert.Metadata.IssueDate = utils.GetTime()
	vp, err := strconv.Atoi(validityPeriod)
	if err != nil {
		fmt.Println("转换失败：", err)
	} else {
		cert.Cert.Metadata.ValidityPeriod = vp
	}
	cert.Cert.Agencies.AgencyName = "江软"
	cert.Cert.Agencies.Address = "0xu2hfuaihvcuwerhuiiwi"
	cert.Cert.Users.UserName = 

}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
