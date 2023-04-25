package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyShow(c *gin.Context) {
	c.HTML(http.StatusOK, "verify.html", nil)
}
