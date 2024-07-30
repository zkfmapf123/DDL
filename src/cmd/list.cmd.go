package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/DDL/src/aws"
)

var listCmd = &cobra.Command{
	Use:   "li",
	Short: "List Lambda Function",
	Long:  "List Lambda Function",
	Run: func(cmd *cobra.Command, args []string) {
		awsConfig, profile, err := aws.New()
		ecrParams := awsConfig.MustBeforeHook()

		if err != nil {
			panic(err)
		}

		fmt.Println(profile, ecrParams)
	},
}

func init() {
	deleteCmd.Flags().String("list", "list", "")
	rootCmd.AddCommand(listCmd)
}
