package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/upload", UploadFile)
	router.GET("/files", ListFiles)
	router.GET("/frequentWords", freqWordsHandler)
	router.GET("/wordcount", wordCountHandler)
	router.DELETE("/files/:filename", DeleteFile)
}
