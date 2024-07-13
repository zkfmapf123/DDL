package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/spf13/viper"
)

type AWSconfig struct {
	LambdaConfig lambda.Client
}

func New() (*AWSconfig, string, error) {
	profile := viper.GetString("profile")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		return &AWSconfig{}, profile, nil
	}

	lambdaConfig := lambda.NewFromConfig(cfg)

	return &AWSconfig{
		LambdaConfig: *lambdaConfig,
	}, profile, nil
}
