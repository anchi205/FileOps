package utils

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dst string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error syncing file")
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error syncing file")
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println("Error syncing file")
	}

	err = destFile.Sync()
	if err != nil {
		fmt.Println("Error syncing file")
	}
}
