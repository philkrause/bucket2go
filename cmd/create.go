package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// options :=

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "create a repository",
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Println("this command will create a repository")
			return nil
		},
	}

	rootCmd.AddCommand(createCmd)
}
