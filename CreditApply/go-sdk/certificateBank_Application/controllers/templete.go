package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PreviewShow(c *gin.Context) {
	c.HTML(http.StatusOK, "preview.html", nil)
}
