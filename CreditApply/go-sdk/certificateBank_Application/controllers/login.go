package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginShow(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
