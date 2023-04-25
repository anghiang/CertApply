package controllers

import (
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CertList(c *gin.Context) {

	var cert models.Cert
	certsSlice, err := cert.QueryCertsByUser()
	if err != nil {
		c.Error(err)
	}
	c.HTML(http.StatusOK, "certsList.html", gin.H{
		"certsSlice": certsSlice,
	})

}
