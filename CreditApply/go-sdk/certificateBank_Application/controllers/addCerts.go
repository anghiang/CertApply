package controllers

import (
	"bytes"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/contracts/implementation"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/services"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/utils"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

var cli *client.Client

var contract *implementation.Contract
var tr *implementation.TransOtp

func AddCertShow(c *gin.Context) {
	session := sessions.Default(c)
	agencyAddr := session.Get(AgencyAddr)

	cli, contract, tr = utils.Init()

	//通过已登录的账号绑定TransOpts和CallOpts
	err := utils.BindOpts(agencyAddr.(string))
	if err != nil {
		fmt.Println("BindOpts: ", err)
		c.JSON(200, err.Error())
		return
	}
	isRole, err := implementation.IsRole(common.HexToAddress(agencyAddr.(string)))
	if err != nil {
		fmt.Println("IsRole : ", err)
		c.JSON(http.StatusBadRequest, "404")
		return
	}
	if !isRole {
		c.Redirect(200, "/login")
		return
	}
	c.HTML(200, "issueCert.html", nil)
}

func AddCertSign(c *gin.Context) {
	validityPeriod := c.PostForm("ValidityPeriod")
	vpInt, err := strconv.Atoi(validityPeriod)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	cert := entity.Certs{
		Cert: models.Cert{
			Metadata: models.Metadata{
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
	session := sessions.Default(c)
	agencyName := session.Get(AgencyName)
	if err != nil {
		fmt.Print("error : ", err)
		return
	}
	cert.Cert.Agencies.AgencyName = agencyName.(string)
	fmt.Println("agencyName : ", agencyName)
	cert.Cert.Agencies.Id, cert.Cert.Agencies.Address, err = cert.Cert.Agencies.QueryAddressByName()
	//通过已登录的账号绑定TransOpts和CallOpts
	err = utils.BindOpts(cert.Cert.Agencies.Address)
	if err != nil {
		fmt.Print("error : ", err)
		c.JSON(200, err.Error())
		return
	}
	cert.Cert.Users.UserName, err = cert.Cert.Users.QueryUNameByAddress()
	if err != nil {
		fmt.Print("error : ", err)
		return
	}
	//计算证书字段的TargetHash
	certTgtHash := cert.TargetHash()
	//_, contract, transOpts := utils.Init()

	tx := contract.Issue(certTgtHash, tr)
	//signTx, err := types.SignTx(tx, types.HomesteadSigner{}, private)
	//	//if err != nil {
	//	//	fmt.Println("SignTx err ", err)
	//	//}
	//	//fmt.Println("SignTx ", signTx)
	//	//r, err := cli.SendTransaction(context.Background(), signTx)
	//	//fmt.Println(r)private, err := crypto.ToECDSA(common.Hex2Bytes("b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62"))
	//
	signature, err := services.SignCertTgtHsh(certTgtHash, cert.Cert.Agencies.Address)
	if err != nil {
		fmt.Println(err)
	}
	cert.Cert.Metadata.Signature = signature
	cert.Cert.AddCert()
	rawTxByte, err := tx.MarshalJSON()
	fmt.Println("rawTxByte: ", rawTxByte)
	if err != nil {
		fmt.Println("tx.MarshalJson err ", err)
		return
	}
	var url = "http://127.0.0.1:5000/setSignature"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(rawTxByte))
	if err != nil {
		// 处理创建请求对象时的错误
		panic(err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer your_token")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// 处理请求发送失败的错误
		panic(err)
	}
	defer resp.Body.Close()
}

func AddCert(c *gin.Context) {

	byteData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println("byteData: ", string(byteData))
	var tx = new(types.Transaction)
	tx.UnmarshalJSON(byteData)
	fmt.Println("tx: ", tx)

	Receipt, err := utils.SendSingedTx(cli, tx)
	if err != nil {
		fmt.Println("Receipt err :", err)
	}
	fmt.Println("Receipt :", Receipt)
	if Receipt.Output != "0x" {
		fmt.Println("该证书已存在")
		c.JSON(http.StatusOK, gin.H{
			"add_status": "add_err",
			"add_msg":    "该证书已存在",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"add_status": "add_ok",
			"add_msg":    "证书颁发成功",
		})
	}
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
