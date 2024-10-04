package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/ddl/internal"
	"github.com/zkfmapf123/ddl/utils"
)

var (
	rootCmd = &cobra.Command{
		Use:   "DDL",
		Short: "Donggyu Driven Lambda",
		Long:  "Donggyu Driven Lambda",
		Run: func(cmd *cobra.Command, args []string) {

			var creds internal.DDLCreds
			t := utils.NewTerminal().Clear()

			for {
				_creds, err := internal.GetAWSCredsProfile()

				// exists ddl creds
				if err == nil {
					creds = _creds
					break
				}

				profiles, err := internal.GetAWSCredsBeforeSaveDDLCreds(utils.AWS_CRED_PATH)
				if err != nil {
					utils.PanicError(err)
				}

				profile := t.SelectOne("Select AWS Profile", profiles)
				region := t.SelectOne("Select AWS Region", utils.AWS_REGIONS)
				err = internal.SetDDLCreds(utils.AWS_DDL_PATH, profile, region)
				if err != nil {
					utils.PanicError(err)
				}
			}

			// Refactoring
			os.Setenv("profile", creds.Profile)
			os.Setenv("region", creds.Region)
		},
	}
)

func initConfig() {

}

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.PanicError(err)
	}
}
