package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/upload", UploadFile)
	router.GET("/files", ListFiles)
	router.DELETE("/files/:filename", DeleteFile)
}
