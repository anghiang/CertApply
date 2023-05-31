package main

import (
	configDb "github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	configDb.Init()
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Static("/assets", "certificateBank_Application/assets")
	r.LoadHTMLGlob("certificateBank_Application/template/*")
	r.GET("/PreviewShow", controllers.PreviewShow)
	r.POST("/VerifyCert", controllers.VerifyCert)
	r.POST("/VerifyTrans", controllers.VerifyTrans)
	//r.GET("/CertToBlock", controllers.GetCertBlockNum)
	r.POST("/download", controllers.DownloadTrans)
	r.POST("/downloadCert", controllers.DownloadCert)
	r.POST("/addCert", controllers.AddCert)
	r.POST("/addCertSign", controllers.AddCertSign)
	r.GET("/addCertShow", controllers.AddCertShow)
	r.GET("/login", controllers.LoginPage)
	r.POST("/addTranscriptSign", controllers.AddTranscriptSign)
	r.GET("/certList", controllers.CertList)
	r.GET("/verify", controllers.VerifyShow)
	r.GET("/revoke", controllers.RevokePage)
	r.POST("/revoke", controllers.RevokeCert)
	r.GET("transcriptPre", controllers.TranscriptPreShow)
	r.POST("/agencyLogin", controllers.AgencyLogin)
	r.POST("/studentLogin", controllers.UserLogin)
	r.POST("/verifySig", controllers.VerifySig)
	r.GET("preview", controllers.PreviewCert)
	r.GET("/certVerifyOk", controllers.CertVerifyOk)
	r.GET("/transVerifyOk", controllers.TransVerifyOk)

	r.Run("localhost:8089")

}
