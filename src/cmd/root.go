package cmd

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zkfmapf123/DDL/src/filesystems"
	"github.com/zkfmapf123/DDL/src/interactions"
)

var (
	DEFAULT_PROFILE = "default"
	CRED_PATH       = ".aws/credentials"
)

var rootCmd = &cobra.Command{
	Use:   "Dobby Driven Lambda",
	Short: "DDL",
	Long:  "Dobby Driven Lambda",
}

func initial() {

	// Cobra onInitialize
	cobra.OnInitialize(func() {

		err := setAWSProfile()
		if err != nil {
			log.Fatalln("[ERROR] set aws profile", err)
		}

		err = setConfig()
		if err != nil {
			log.Fatalln("[ERROR] init config", err)
		}
	})

	// Bind Flags
	rootCmd.PersistentFlags().StringP("profile", "p", "", "[Required] AWS Profile")
	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))
}

func setConfig() error {

	if viper.GetString("profile") == "" {
		viper.Set("profile", DEFAULT_PROFILE)
	}

	return nil
}

func setAWSProfile() error {
	homePath := filesystems.MustGetHomePath()
	joinPath := filepath.Join(homePath, CRED_PATH)

	// credentials 파일은 init 추가
	viper.SetConfigFile(joinPath)
	viper.SetConfigType("ini")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// region
	region := interactions.Prompt("Input Your Region : ")
	if region == "" {
		return errors.New("InValid Region")
	}

	viper.Set("region", region)
	return nil
}

func Execute() error {
	initial()

	return rootCmd.Execute()
}
