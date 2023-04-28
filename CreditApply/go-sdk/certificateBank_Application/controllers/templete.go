package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PreviewShow(c *gin.Context) {
	c.HTML(http.StatusOK, "preview.html", nil)
}

func TranscriptPreShow(c *gin.Context) {
	c.HTML(http.StatusOK, "transcriptPre.html", nil)
}
