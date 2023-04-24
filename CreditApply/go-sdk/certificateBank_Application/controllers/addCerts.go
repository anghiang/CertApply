package controllers

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
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
	validityPeriod := c.PostForm("ValidityPeriod")
	vpInt, err := strconv.Atoi(validityPeriod)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	cert := entity.Certs{
		Cert: models.Cert{
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
	//agencyName, _ := c.Request.Cookie("agencyName")
	cert.Cert.Agencies.AgencyName = "清华大学"
	//cert.Cert.Agencies.Address, err = cert.Cert.Agencies.QueryAddressByName()
	cert.Cert.Agencies.Address = "0x83309d045a19c44dc3722d15a6abd472f95866ac"
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"add_status": "add_err",
		})
		fmt.Print("error : ", err)
		return
	}
	cert.Cert.Users.UserName, err = cert.Cert.Users.QueryUNameByAddress()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"add_status": "add_err",
		})
		fmt.Print("error : ", err)
		return
	}
	certTgtHash := cert.TargetHash()
	err = services.IssueCert(cert, certTgtHash)
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

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
