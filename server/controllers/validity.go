package controllers

import (
	"net/http"

	"github.com/anchi205/FileOps/server/store"
	"github.com/gin-gonic/gin"
)

func GetValidityHandler(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	file_hash := c.Query("file_hash")
	file_name := c.Query("file_name")
	if _, ok := store.Hashstore[file_hash]; ok {
		for _, name := range store.Hashstore[file_hash] {
			if name == file_name {
				c.JSON(http.StatusOK, gin.H{"validity": []bool{true, true}})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"validity": []bool{true, false}})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"validity": []bool{false, false}})
		return
	}
}
