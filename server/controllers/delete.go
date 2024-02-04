package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DeleteFile(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	filename := c.PostForm("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not specified"})
		return
	}

	filePath := filepath.Join("./uploads", filename)
	if err := handleSpecifiedFileDeletion(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateFileHash(filename)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File %s removed successfully!", filename)})
}

func handleSpecifiedFileDeletion(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func updateFileHash(filename string) {
	hash, _ := calculateFileHash(filepath.Join("./uploads", filename))
	FileHash[hash] = append(FileHash[hash], filename)
}
