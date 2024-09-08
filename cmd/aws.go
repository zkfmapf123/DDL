package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/samber/lo"
)

type awsOpts struct {
	profile      string
	awsAccessKey string
	awsSecretKey string
	region       string
}

func NewAWSCredentials() *awsOpts {

	return &awsOpts{
		region: "ap-northeast-2",
	}
}

func (opt *awsOpts) WithRegion(region string) *awsOpts {
	opt.region = region
	return opt
}

func (opt *awsOpts) WithAccessSecretKey(accessKey, secretKey string) *awsOpts {
	opt.awsAccessKey = accessKey
	opt.awsSecretKey = secretKey
	return opt
}

func (opt *awsOpts) WithProfile(profile string) *awsOpts {
	opt.profile = profile
	return opt
}

type AWSClientParams struct {
	SSMConfig    ssm.Client
	LambdaConfig lambda.Client
	ECRConfig    ecr.Client
	IAMConfig    iam.Client
}

func (opt *awsOpts) MustEnd() AWSClientParams {

	var client aws.Config
	var err error

	// proifle이 있는 경우
	if !lo.IsEmpty(opt.profile) {
		client, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithSharedConfigProfile(opt.profile),
			config.WithRegion(opt.region),
		)

		if err != nil {
			panic(err)
		}

	} else {

		creds := credentials.NewStaticCredentialsProvider(opt.awsAccessKey, opt.awsSecretKey, "")

		// accessKey, secretKey로 구성하는 경우
		client, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithCredentialsProvider(creds),
		)

		if err != nil {
			panic(err)
		}

	}

	return AWSClientParams{
		SSMConfig:    *ssm.NewFromConfig(client),
		LambdaConfig: *lambda.NewFromConfig(client),
		ECRConfig:    *ecr.NewFromConfig(client),
		IAMConfig:    *iam.NewFromConfig(client),
	}
}
