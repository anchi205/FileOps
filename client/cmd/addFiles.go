package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/anchi205/FileOps/client/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

type ValidityResponse struct {
	Validity []bool `json:"validity"`
}

func getValidityFromServer(file_hash string, file_name string) []bool {
	url := "http://localhost:8080/validity?file_hash=" + file_hash + "&file_name=" + file_name
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while pinging server: ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	var response ValidityResponse
	// Unmarshal the JSON string into the ValidityResponse struct
	if err := json.Unmarshal([]byte(string(body)), &response); err != nil {
		fmt.Println("Error in parsing the response", err)
	}
	return response.Validity
}

func uploadFile(url, filePath, fileName, fileHash string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("filename", fileName)
	writer.WriteField("filehash", fileHash)

	// Create a form file field and add it to the multipart writer
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return err
	}

	// Copy the file data to the part
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

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
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}

func create_duplicate_file(url, fileName, fileHash string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
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
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}

func handleAddFiles(args []string) {
	// args has all the file paths
	for _, file := range args {
		file_content := utils.ReadFile(file)
		file_hash, err := utils.GetSHA1(file_content)
		if err != nil {
			fmt.Println("Error while getting SHA1 hash of file: ", file)
		}
		file_name := utils.TrimFileName(file)
		validity := getValidityFromServer(file_hash, file_name)
		if validity[0] {
			if !validity[1] {
				create_duplicate_file("http://localhost:8080/createDuplicate", file_name, file_hash)
			} else {
				fmt.Println("File already exists on server, no need to upload again")
			}
		} else {
			uploadFile("http://localhost:8080/upload", file, file_name, file_hash)
		}
	}
}

func addFilesCLIHandler(args []string) {
	handleAddFiles(args)
}

// ----on cli----
// get multiple files from cli and process them by making file_hash, file_content and file_name
// check for hash on the server if it exists then check for file name (make a check_presence() endpoint on server for this)
// if both exists do nothing
// if only hash exists then just create a new file on server
// if none exists then upload the file and create a new file on server

// ----on server----
// mape of hash and []file_name
// example : 897ew7d647ds6876ds => [file1, file2, file3]

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add files command",
	Long:  "add files command",
	Run: func(cmd *cobra.Command, args []string) {
		addFilesCLIHandler(args)
	},
}
