package controllers

import (
	"fmt"

	"github.com/anchi205/FileOps/client/utils"
	"github.com/anchi205/FileOps/server/store"
	"github.com/gin-gonic/gin"
)

func UploadFileHandler(c *gin.Context) {

	file, err := c.FormFile("file")
	fileName := c.PostForm("filename")
	fileHash := c.PostForm("filehash")

	if err != nil {
		c.JSON(400, gin.H{"error": "File not found in request"})
		return
	}

	// Save the file locally
	if err := c.SaveUploadedFile(file, "uploads/"+fileName); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}
	fmt.Printf("Uploaded file: %s, Filehash: %s\n", fileName, fileHash)
	store.Hashstore[fileHash] = []string{fileName}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("File %s uploaded successfully", fileName),
	})
}

func CreateDuplicateHandler(c *gin.Context) {
	fileName := c.PostForm("filename")
	fileHash := c.PostForm("filehash")

	store.Hashstore[fileHash] = append(store.Hashstore[fileHash], fileName)
	utils.CopyFile("uploads/"+store.Hashstore[fileHash][0], "uploads/"+fileName)

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Duplicate file %s created successfully", fileName),
	})
}
