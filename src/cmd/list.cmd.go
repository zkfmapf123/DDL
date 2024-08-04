package cmd

import (
	"errors"
	"fmt"
	"reflect"

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

		if profile == "" || ecrParams.Name == "" {
			panic(errors.New("no Values"))
		}

		res, err := awsConfig.ListLambda()
		if err != nil {
			panic(err)
		}

		t := reflect.TypeOf(res[0])
		for i := 0; i < t.NumField(); i++ {
			fmt.Println(t.Field(i).Name)
		}

		for k, v := range res {
			fmt.Println(k, v)
		}
	},
}

func init() {
	deleteCmd.Flags().String("list", "list", "")
	rootCmd.AddCommand(listCmd)
}
