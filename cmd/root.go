package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bucket2go",
	Short: "This CLI is used for creating and managing repositories in bitbucket",
	Long:  "This CLI is used for creating and managing repositories in bitbucket",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
