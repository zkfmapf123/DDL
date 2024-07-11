package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "Dobby Driven Lambda",
	Short: "DDL",
	Long:  "Dobby Driven Lambda",
}

func initial() {

}

func Execute() error {
	initial()

	return rootCmd.Execute()
}
