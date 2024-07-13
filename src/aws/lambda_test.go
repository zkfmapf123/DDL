package aws

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

var lambdaInfo = LambdaParams{
	Name: "hello_function",
}

func Test_retrive(t *testing.T) {

	cfg, _, _ := New()
	viper.Set("profile", "root")

	output, _ := retriveLambda(cfg.LambdaConfig, lambdaInfo)
	fmt.Println(output)
}
