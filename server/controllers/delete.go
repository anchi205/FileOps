package controllers

import (
	"fmt"
	"os"

	"github.com/anchi205/FileOps/server/store"
	"github.com/gin-gonic/gin"
)

func DeleteFileHandler(c *gin.Context) {

	fileName := c.PostForm("filename")
	fileHash := c.PostForm("filehash")

	// delete fileName from disk
	fmt.Println("Deleting file:", fileName)
	fileName = "uploads/" + fileName
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	// delete fileName from store
	if len(store.Hashstore[fileHash]) == 1 {
		delete(store.Hashstore, fileHash)
		return
	}

	for i, name := range store.Hashstore[fileHash] {
		if name == fileName {
			store.Hashstore[fileHash] = append(store.Hashstore[fileHash][:i], store.Hashstore[fileHash][i+1:]...)
		}
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Deleted file %s  successfully", fileName),
	})
}
