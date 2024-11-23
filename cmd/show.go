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
			creds, err := internal.GetAWSCredsProfile()
			if err != nil {
				utils.PanicError(err)
			}

			awsCreds := internal.NewAWSCredentials().
				WithProfile(creds.Profile).
				WithRegion(creds.Region).
				MustEnd()

			showCreds(creds, *table)
			showSSM(awsCreds, *table)
			showLambda(awsCreds, *table)
		},
	}
)

// show aws credentials
func showCreds(creds internal.DDLCreds, table utils.TwParams) {
	table.Setheader([]string{"profile", "region"})
	table.SetBody([][]string{{creds.Profile, creds.Region}})
	table.Print()
}

// show ssm key / value
func showSSM(awsCreds internal.AWSClientParams, table utils.TwParams) {
	key, value, err := awsCreds.GetSSMKeyValue()
	if err != nil {
		utils.PanicError(err)
	}

	table.Setheader([]string{"ssm-Key", "ssm-value"})
	table.SetBody([][]string{{key, value}})
	table.Print()
}

// show lamba use ddl
func showLambda(awsCreds internal.AWSClientParams, table utils.TwParams) {

}

func init() {
	rootCmd.AddCommand(showCmd)
}
