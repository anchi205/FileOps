package routers

import (
	"github.com/anchi205/FileOps/server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/upload", controllers.UploadFileHandler)
	router.POST("/createDuplicate", controllers.CreateDuplicateHandler)
	router.GET("/list", controllers.ListFiles)
	router.GET("/validity", controllers.GetValidityHandler)
	router.GET("/frequentWords", controllers.FreqWordsHandler)
	router.GET("/wordcount", controllers.WordCountHandler)
	router.POST("/delete", controllers.DeleteFileHandler)
}
