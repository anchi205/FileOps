package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

func removeCLIHandler(args []string) {
	fileName := args[0]
	url := os.Getenv("BASE_URL") + "/removefile"
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Failed to create HTTP request.", err)
		return
	}

	query := request.URL.Query()
	query.Add("filename", fileName)
	request.URL.RawQuery = query.Encode()
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Failed to send HTTP request:", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("API request failed. Status: ", response.Status)
		return
	}

	fmt.Println("File removed successfully !")
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove file command",
	Long:  "remove file command",
	Run: func(cmd *cobra.Command, args []string) {
		removeCLIHandler(args)
	},
}
