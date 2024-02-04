package main

import (
	"github.com/anchi205/FileOps/server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up routes for file operations
	controllers.SetupRoutes(router)

	// Set up routes for hash-related operations
	controllers.SetupHashRoutes(router)

	router.Run(":8080")
}
