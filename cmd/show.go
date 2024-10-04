package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zkfmapf123/ddl/utils"
)

var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show Info DDL",
		Long:  "Show Info DDL",
		Run: func(cmd *cobra.Command, args []string) {
			utils.NewTerminal().Clear()

			table := utils.NewTableWriter()
		},
	}
)

func showCreds(table utils.TwParams) {

}

func showSSM() {

}

func showLambda() {

}

func init() {

	rootCmd.AddCommand(showCmd)
}
