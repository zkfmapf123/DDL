package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "cr",
	Short: "Create Lambda Function",
	Long:  "Create Lambda Function",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Lambda Function")
	},
}

func init() {
	createCmd.Flags().String("create", "create", "")
	rootCmd.AddCommand(createCmd)
}
