package cmd

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const serverURL = "http://localhost:8080/store/add"

// ensures that the addCmd is available as a subcommand of the main command.
func init() {
	rootCmd.AddCommand(addCmd)
}

func calculateFileHash(filePath string) (string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	hasher := sha256.New()
	hasher.Write(fileData)
	hash := hex.EncodeToString(hasher.Sum(nil)) // hexadecimal repr

	return hash, nil
}

func createFileHashJSON(args []string) []byte {
	var FileHash = make(map[string]string)
	for _, file := range args {
		hash, err := calculateFileHash(file)
		if err != nil {
			fmt.Println("Error calculating hash for file:", err)
		}
		FileHash[file] = hash
	}
	jsonFileHashData, err := json.Marshal(FileHash)
	if err != nil {
		fmt.Println(err)
	}
	return jsonFileHashData
}

func filterFilesToUpload(jsonFileHashData []byte) map[string]interface{} {
	baseURL := os.Getenv("BASE_URL")
	url := baseURL + "/getHash"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonFileHashData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	filteredFiles := make(map[string]interface{})
	json.Unmarshal([]byte(responseBody), &filteredFiles)

	return filteredFiles
}

func uploadFilesToServer(files map[string]interface{}) {
	baseURL := os.Getenv("BASE_URL")
	url := baseURL + "/upload"

	for key, value := range files {
		// key is file name and value is isPresent boolean
		file, err := os.Open(key)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		parts := strings.Split(key, "/")
		parsedFileName := parts[len(parts)-1]
		hash, _ := calculateFileHash(key)

		if value == true {
			_ = writer.WriteField("isFilePresent", "true")
		} else {
			_ = writer.WriteField("isFilePresent", "false")
		}
		_ = writer.WriteField("fileName", parsedFileName)
		_ = writer.WriteField("fileHash", hash)

		formFile, err := writer.CreateFormFile("files", parsedFileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(formFile, file)
		writer.Close()

		req, err := http.NewRequest("POST", url, body)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(bodyBytes))
	}
}

func addFilesCLIHandler(args []string) {
	jsonFileHashData := createFileHashJSON(args)
	filteredFiles := filterFilesToUpload(jsonFileHashData)
	uploadFilesToServer(filteredFiles)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "updates files command",
	Long:  "updates files command",
	Run: func(cmd *cobra.Command, args []string) {
		addFilesCLIHandler(args)
	},
}
