package server

import (
	"github.com/anchi205/FileOps/server/routers"
	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()

	// Set up routes for file operations
	routers.SetupRoutes(router)

	// Set up routes for hash-related operations
	// controllers.SetupHashRoutes(router)

	router.Run(":8080")
}
