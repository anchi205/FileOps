package routers

import (
	"github.com/anchi205/FileOps/server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/upload", controllers.UploadFile)
	router.GET("/files", controllers.ListFiles)
	router.GET("/frequentWords", controllers.FreqWordsHandler)
	router.GET("/wordcount", controllers.WordCountHandler)
	router.DELETE("/files/:filename", controllers.DeleteFile)
}
