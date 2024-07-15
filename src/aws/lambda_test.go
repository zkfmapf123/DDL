package aws

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var lambdaInfo = LambdaParams{
	Name: "hello_function",
}

func beforeAll() *AWSconfig {

	viper.Set("profile", "root")
	viper.Set("region", "ap-northeast-2")
	cfg, _, _ := New()
	return cfg
}

func Test_retrive(t *testing.T) {

	cfg := beforeAll()
	output, _ := retriveLambda(cfg.LambdaConfig, lambdaInfo)

	assert.Equal(t, output.Name == "", false)
	assert.Equal(t, output.Arn == "", false)
}

func Test_retrives(t *testing.T) {
	cfg := beforeAll()
	output, err := retriveLambdas(cfg.LambdaConfig)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, len(output) == 0, false)
}
