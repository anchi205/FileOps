package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateFileCLIHandler(args []string) {
	jsonFileHashData := createFileHashJSON(args)
	extractedFiles := filterFilesToUpload(jsonFileHashData)
	uploadFilesToServer(extractedFiles)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updating file cmd",
	Long:  "updating file cmd",
	Run: func(cmd *cobra.Command, args []string) {
		updateFileCLIHandler(args)
	},
}
