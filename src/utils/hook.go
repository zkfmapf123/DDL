package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func BeforeHook() {
	fmt.Println("Profile :", viper.GetString("profile"))
	fmt.Println("Region :", viper.GetString("region"))
}
