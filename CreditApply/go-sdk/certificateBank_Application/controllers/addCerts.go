package controllers

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/services"
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
	cert.Cert.Agencies.AgencyName = "清华"
	cert.Cert.Agencies.Address = "0xu2hfuaihvcuwerhuiiwi"

	uname, err := models.QueryUNameByAddress(address)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"add_staute": "add_err",
		})
		panic(err)

	}
	cert.Cert.Users.UserName = uname
	cert.Cert.Users.Address = address
	cert.Cert.Metadata.Signature = "qwvijnwijnvijbwivhjbihjwbqivbijwq"

	//cert.Cert.Metadata.TargetHash =
	err = services.IssueCert(cert)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"add_staute": "add_ok",
	})
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
