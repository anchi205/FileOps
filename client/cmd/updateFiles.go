package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateFileCLIHandler(args []string) {
	handleAddFiles(args)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update file command",
	Long:  "update file command",
	Run: func(cmd *cobra.Command, args []string) {
		updateFileCLIHandler(args)
	},
}
