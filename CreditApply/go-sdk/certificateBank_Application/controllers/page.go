package controllers

import (
	"errors"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CertList(c *gin.Context) {
	session := sessions.Default(c)
	userAddr := session.Get(UserAddr)
	if userAddr.(string) == "" {
		c.Error(errors.New("未登录的用户"))
		return
	}
	var cert models.Cert
	cert.Users.Address = userAddr.(string)
	certsSlice, err := cert.QueryCertsByUser()
	if err != nil {
		c.Error(err)
	}
	var trans models.Transcript
	trans.Users.Address = userAddr.(string)
	TranscriptSlice, err := trans.QueryTranscriptByUser()
	if err != nil {
		c.Error(err)
	}
	fmt.Println(certsSlice)
	c.HTML(200, "certsList.html", gin.H{
		"certsSlice":      certsSlice,
		"TranscriptSlice": TranscriptSlice,
	})

}
func PreviewCert(c *gin.Context) {

	session := sessions.Default(c)
	userAddress := session.Get(UserAddr)
	if userAddress.(string) == "" {
		c.Error(errors.New("未登录的用户"))
		return
	}
	certId := c.Query("cert_id")
	certType := c.Query("cert_type")
	fmt.Println(certId)

	if certType == "1" {
		var cert models.Cert
		cert.Users.Address = userAddress.(string)
		cert.Metadata.Id, _ = strconv.Atoi(certId)
		cert.CertType, _ = strconv.Atoi(certType)
		err := cert.QueryCertsById()
		if err != nil {
			fmt.Println(err)
		}
		c.HTML(200, "certPre.html", gin.H{
			"certificate": cert,
		})
		return
	} else if certType == "2" {
		var trans models.Transcript
		trans.Users.Address = userAddress.(string)
		trans.Metadata.Id, _ = strconv.Atoi(certId)
		trans.CertType, _ = strconv.Atoi(certType)
		err := trans.QueryTranscriptsById()
		fmt.Println(trans.Course)
		if err != nil {
			fmt.Println(err)
		}
		c.HTML(200, "preview.html", gin.H{
			"certificate": trans,
		})
		return
	}
	return
}

func RevokePage(c *gin.Context) {
	session := sessions.Default(c)
	agencyAddr := session.Get(AgencyAddr)
	if agencyAddr.(string) == "" {
		c.Error(errors.New("未登录的用户"))
		return
	}
	c.HTML(200, "revoke.html", nil)
}

func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func PreviewShow(c *gin.Context) {
	c.HTML(200, "preview.html", nil)
}

func TranscriptPreShow(c *gin.Context) {
	c.HTML(200, "certPre.html", nil)
}

func VerifyShow(c *gin.Context) {
	c.HTML(200, "verify.html", nil)
}

func CertVerifyOk(c *gin.Context) {
	session := sessions.Default(c)
	cert := session.Get("certificate").(*entity.Certs)

	if cert == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "certVerifyOk.html", gin.H{
		"certificate": cert,
	})
}

func TransVerifyOk(c *gin.Context) {
	session := sessions.Default(c)
	Trans := session.Get("certificate").(*entity.TranscriptCert)

	var isHiddenData bool

	if len(Trans.HiddenData) == 0 {
		isHiddenData = false
	} else {
		isHiddenData = true
	}
	if Trans == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(isHiddenData)
	c.HTML(http.StatusOK, "TransVerifyOk.html", gin.H{
		"certificate": Trans,
		"isHidden":    isHiddenData,
	})
}
