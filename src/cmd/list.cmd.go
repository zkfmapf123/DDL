package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/DDL/src/aws"
	"github.com/zkfmapf123/DDL/src/utils"
)

var listCmd = &cobra.Command{
	Use:   "li",
	Short: "List Lambda Function",
	Long:  "List Lambda Function",
	Run: func(cmd *cobra.Command, args []string) {
		utils.BeforeHook()
		awsConfig, profile, err := aws.New()

		if err != nil {
			panic(err)
		}

		fmt.Println(awsConfig, profile)
	},
}

func init() {
	deleteCmd.Flags().String("list", "list", "")
	rootCmd.AddCommand(listCmd)
}
