package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/DDL/src/aws"
)

var createCmd = &cobra.Command{
	Use:   "cr",
	Short: "Create Lambda Function",
	Long:  "Create Lambda Function",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _, err := aws.New()
		if err != nil {
			panic(err)
		}

		fmt.Println(cfg)

	},
}

func init() {
	createCmd.Flags().String("create", "create", "")
	rootCmd.AddCommand(createCmd)
}
