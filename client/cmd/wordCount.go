package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(wcCmd)
	wcCmd.Flags().StringP("order", "s", "dsc", "sort order")
	wcCmd.Flags().StringP("limit", "l", "10", "limit")
}

type wcResponse struct {
	WordCount int `json:"word_count"`
}

var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "word count command",
	Long:  "word count command",
	Run: func(cmd *cobra.Command, args []string) {
		sortorder, _ := cmd.Flags().GetString("order")
		limit, _ := cmd.Flags().GetString("limit")
		baseURL := "http://localhost:8080/" + "wordcount"
		if sortorder != "" {
			baseURL = baseURL + "?sortOrder=" + sortorder
		}
		if limit != "" {
			separator := "?"
			if sortorder != "" {
				separator = "&"
			}
			baseURL = baseURL + separator + "limit=" + limit
		}
		wordCountCLIHandler(baseURL)
	},
}

func wordCountCLIHandler(baseURL string) {
	response, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var res wcResponse
	if err := json.Unmarshal([]byte(string(body)), &res); err != nil {
		fmt.Println("Error in parsing the response", err)
	}
	message := "Total number of words in all files : " + fmt.Sprint(res.WordCount)
	fmt.Println(message)
}
