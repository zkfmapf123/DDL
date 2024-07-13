package cmd

import (
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zkfmapf123/DDL/src/filesystems"
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

	return nil
}

func Execute() error {
	initial()

	return rootCmd.Execute()
}
