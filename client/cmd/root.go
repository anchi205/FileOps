package cmd

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fileops",
	Short: "FileOps is a CLI",
	Long:  "FileOps is a CLI",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func Execute() error {
	godotenv.Load()
	return rootCmd.Execute()
}
