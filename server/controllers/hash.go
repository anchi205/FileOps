package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func calculateFileHash(filePath string) (string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(fileData)
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash, nil
}

func getHashHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	getIsFileHashed(c)
}

func getIsFileHashed(c *gin.Context) {
	var input map[string]string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output := make(map[string]bool)
	for key, value := range input {
		hash, err := calculateFileHash(value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		_, ok := FileHash[hash]
		output[key] = ok
	}

	c.JSON(http.StatusOK, output)
}

func SetupHashRoutes(router *gin.Engine) {
	router.POST("/getHash", getHashHandler)
}
