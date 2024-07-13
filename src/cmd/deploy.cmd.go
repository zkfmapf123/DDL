package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "dep",
	Short: "Deploy Lambda Function",
	Long:  "Deploy Lambda Function",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deploy Lambda Function")
	},
}

func init() {
	deleteCmd.Flags().String("deploy", "deploy", "")
	rootCmd.AddCommand(deployCmd)
}
