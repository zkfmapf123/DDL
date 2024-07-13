package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "de",
	Short: "Delete Lambda Function",
	Long:  "Delete Lambda Function",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete Lambda Function")
	},
}

func init() {
	deleteCmd.Flags().String("delete", "delete", "")
	rootCmd.AddCommand(deleteCmd)
}
