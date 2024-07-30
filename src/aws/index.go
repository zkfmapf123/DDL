package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/spf13/viper"
)

type AWSconfig struct {
	LambdaConfig lambda.Client
	ECRConfig    ecr.Client
	SSMConfig    ssm.Client
}

func New() (*AWSconfig, string, error) {
	profile := viper.GetString("profile")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		return &AWSconfig{}, profile, nil
	}

	lambdaConfig := GetLambdaConfig(cfg)
	ecrConfig := GetECRConfig(cfg)
	ssmConfig := GetSSMConfig(cfg)

	return &AWSconfig{
		LambdaConfig: *lambdaConfig,
		ECRConfig:    *ecrConfig,
		SSMConfig:    *ssmConfig,
	}, profile, nil
}
