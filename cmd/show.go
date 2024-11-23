package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zkfmapf123/ddl/internal"
	"github.com/zkfmapf123/ddl/utils"
)

var (
	showCmd = &cobra.Command{
		Use:   "s",
		Short: "Show Info DDL",
		Long:  "Show Info DDL",
		Run: func(cmd *cobra.Command, args []string) {
			utils.NewTerminal().Clear()

			table := utils.NewTableWriter()

			showCreds(*table)
			showSSM(*table)
			showLambda(*table)
		},
	}
)

// show aws credentials
func showCreds(table utils.TwParams) {
	creds, err := internal.GetAWSCredsProfile()
	if err != nil {
		utils.PanicError(err)
	}

	table.Setheader([]string{"profile", "region"})
	table.SetBody([][]string{{creds.Profile, creds.Region}})
	table.Print()
}

// show ssm key / value
func showSSM(table utils.TwParams) {

}

// show lamba use ddl
func showLambda(table utils.TwParams) {

}

func init() {
	rootCmd.AddCommand(showCmd)
}
