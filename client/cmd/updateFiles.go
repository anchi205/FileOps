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
	Short: "update file command",
	Long:  "update file command",
	Run: func(cmd *cobra.Command, args []string) {
		updateFileCLIHandler(args)
	},
}
