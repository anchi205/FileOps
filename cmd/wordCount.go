package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(wcCmd)
	wcCmd.Flags().StringP("order", "s", "dsc", "sort order")
	wcCmd.Flags().StringP("limit", "l", "10", "limit")
}

var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "word count command",
	Long:  "word count command",
	Run: func(cmd *cobra.Command, args []string) {
		sortorder, _ := cmd.Flags().GetString("order")
		limit, _ := cmd.Flags().GetString("limit")
		baseURL := os.Getenv("BASE_URL") + "/wordcount"
		wordCountCLIHandler(baseURL)
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

	fmt.Println(string(body))
}
