package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/anchi205/FileOps/client/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

func delete_file(url, fileName, fileHash string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error while deleting file: ", err)
	}

	writer.WriteField("filename", fileName)
	writer.WriteField("filehash", fileHash)

	// Close the multipart writer
	writer.Close()

	// Set the content type header
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Set the request body
	req.Body = io.NopCloser(body)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while deleting file: ", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error while deleting file: ", err)
	}
}

func removeCLIHandler(args []string) {
	fileName := args[0]
	file_content := utils.ReadFile(fileName)
	file_hash, err := utils.GetSHA1(file_content)
	if err != nil {
		fmt.Println("Error while getting SHA1 hash of file: ", fileName)
	}
	file_name := utils.TrimFileName(fileName)
	delete_file("http://localhost:8080/delete", file_name, file_hash)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove file command",
	Long:  "remove file command",
	Run: func(cmd *cobra.Command, args []string) {
		removeCLIHandler(args)
	},
}
