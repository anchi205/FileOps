package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	isFilePresent := c.PostForm("isFilePresent") == "true"
	fileName := c.PostForm("fileName")
	fileHash := c.PostForm("fileHash")

	if isFilePresent {
		if _, err := os.Stat(filepath.Join("./uploads", fileName)); err == nil {
			c.String(http.StatusOK, "File already present on the server!")
			return
		} else {
			// file does not exist
			c.String(http.StatusOK, "File content already present on the server!, creating new file")
			srcFile := FileHash[fileHash][0]
			source, err := os.Open(filepath.Join("./uploads", srcFile))
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			defer source.Close()

			dst, err := os.Create(filepath.Join("./uploads", fileName))
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, source)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			FileHash[fileHash] = append(FileHash[fileHash], fileName)
			return
		}
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	files := form.File["files"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer file.Close()

		dst, err := os.Create(filepath.Join("./uploads", fileHeader.Filename))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		hash, err := calculateFileHash(filepath.Join("./uploads", fileHeader.Filename))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		FileHash[hash] = append(FileHash[hash], fileHeader.Filename)
	}

	message := fmt.Sprintf("%s uploaded successfully!", fileName)
	c.String(http.StatusOK, message)
}
