package controllers

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListFiles(c *gin.Context) {
	dir := "./uploads"

	files, err := os.ReadDir(dir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}

	response := strings.Join(fileNames, "\n")
	c.String(http.StatusOK, response)
}
