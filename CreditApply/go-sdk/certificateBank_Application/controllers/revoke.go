package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RevokeShow(c *gin.Context) {
	c.HTML(http.StatusOK, "revoke.html", nil)
}
