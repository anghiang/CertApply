package controllers

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const AgencyName = "agencyName"
const AgencyAddr = "agencyAddr"
const UserAddr = "userAddr"

func AgencyLogin(c *gin.Context) {
	agency := models.Agency{
		Address: c.PostForm("AgencyAddress"),
	}
	pwd := c.PostForm("Password")
	var err error
	agency.AgencyName, err = agency.Login(pwd)
	fmt.Println(agency.AgencyName)
	if err != nil {
		c.JSON(200, gin.H{
			"login_status": "login_err",
		})
		return
	}
	session := sessions.Default(c)
	session.Set(AgencyName, agency.AgencyName)
	session.Set(AgencyAddr, agency.Address)
	session.Save()
	c.JSON(200, gin.H{
		"login_status": "login_ok",
	})
}

func UserLogin(c *gin.Context) {
	user := models.User{
		Address: c.PostForm("StudentAddress"),
	}
	pwd := c.PostForm("Password")
	var err error
	user.UserName, err = user.Login(pwd)
	if err != nil {
		c.JSON(200, gin.H{
			"login_status": "login_err",
		})
		return
	}

	session := sessions.Default(c)
	session.Set(UserAddr, user.Address)
	session.Save()
	c.JSON(200, gin.H{
		"login_status": "login_ok",
	})
}
