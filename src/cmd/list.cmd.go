package cmd

import (
	"errors"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/DDL/src/aws"
	"github.com/zkfmapf123/DDL/src/utils"
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

		if profile == "" || ecrParams.Name == "" {
			panic(errors.New("no Values"))
		}

		res, err := awsConfig.ListLambda()
		if err != nil {
			panic(err)
		}

		t := reflect.TypeOf(res[0])
		headers := []string{}
		for i := 0; i < t.NumField(); i++ {
			headers = append(headers, t.Field(i).Name)
		}

		utils.PrintDashboardUseKey(headers, res)
	},
}

func init() {
	deleteCmd.Flags().String("list", "list", "")
	rootCmd.AddCommand(listCmd)
}
