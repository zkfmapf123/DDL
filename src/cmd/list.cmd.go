package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "li",
	Short: "List Lambda Function",
	Long:  "List Lambda Function",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List Lambda Function")
	},
}

func init() {
	deleteCmd.Flags().String("list", "list", "")
	rootCmd.AddCommand(listCmd)
}
