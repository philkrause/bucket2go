package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {

	var timeCmd = &cobra.Command{
		Use:   "time",
		Short: "time check",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(time.Now())
		},
	}

	rootCmd.AddCommand(timeCmd)
}
