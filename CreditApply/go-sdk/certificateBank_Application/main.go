package main

import (
	configDb "github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/contracts/implementation"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/controllers"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}
	//implementation.InitContract(client)
	//var czc common.Address = common.HexToAddress("0x83309d045a19c44dc3722d15a6abd472f95866ac")
	//
	//h := implementation.GetCert(czc)
	//for _, certificate := range h {
	//	fmt.Println("cert : ", certificate)
	//}
	//priByte, err := utils.GetPrivateKey()
	//if err != nil {
	//	fmt.Println("err :", err)
	//}
	//fmt.Println("priByte", priByte)
	//config.Init()
	//c := new(models.Transcript)
	//
	//c.Metadata.Name = "2019学年期末成绩单"
	//c.Metadata.Number = "2131231"
	//c.Metadata.IssueDate = "2019-6-30"
	//c.Metadata.ValidityPeriod = 100
	//c.Metadata.Signature = "0xuywe9ihdhcjkzhdAJKy1u93y897374983"
	//c.Agencies.Id = 1
	//c.Users.Id = 1
	//cour1 := new(models.Course)
	//cour2 := new(models.Course)
	//
	//cour1.CourseName = "java"
	//cour1.Score = 100
	//cour2.CourseName = "web"
	//cour2.Score = 90
	//c.Course = append(c.Course, *cour1)
	//c.Course = append(c.Course, *cour2)
	////err := c.AddTranscript("0x82943A2985335485F11405008fcBa4A0b9e5096c")
	//fmt.Println(c.QueryTranscriptsByUser("0x82943A2985335485F11405008fcBa4A0b9e5096c"))
	//res, err := c.QueryCertByUser("0x82943A2985335485F11405008fcBa4A0b9e5096c")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("cert := ", res)
	//var t *models.Transcript
	//t.Metadata.Id = 1
	//t.Metadata.Name = "2019学年期末成绩单"
	//t.Metadata.Number = "2131231"
	//t.Metadata.IssueDate = "2019-6-30"
	//t.Metadata.ValidityPeriod = 100
	//t.Metadata.Signature = "0xuywe9ihdhcjkzhdAJKy1u93y897374983"
	//t.Metadata.TargetHash = "0x98740w6r78qtdguiasbhfqgusobuadhks"
	//t.Users.Id = 1
	//t.Agencies.Id = 1
	//t.Course = append(t.Course)
	//err := t.AddTranscript()
	//entity.CalHash()
	implementation.InitContract(client)
	configDb.Init()
	r := gin.Default()
	r.Static("/assets", "certificateBank_Application/assets")
	r.LoadHTMLGlob("certificateBank_Application/template/*")
	r.GET("/PreviewShow", controllers.PreviewShow)
	r.POST("/VerifyCert", controllers.VerifyCert)
	r.GET("/CertToBlock", controllers.GetCertBlockNum)
	r.POST("/download", controllers.Download)
	r.POST("/addCert", controllers.AddCert)
	r.GET("/addCertShow", controllers.AddCertShow)
	r.GET("/login", controllers.LoginShow)
	r.POST("/addTranscript", controllers.AddTranscript)
	r.GET("/certList", controllers.CertList)
	r.GET("/verify", controllers.VerifyShow)
	r.GET("revoke", controllers.RevokeShow)
	r.Run("localhost:8089")

}
