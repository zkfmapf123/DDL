package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

const (
	SSM_PATH = "/ddl/store/ecr"
	ECR_NAME = "ddl-repo"
)

/*
* Provider : AWS
 */
func GetSSMConfig(cfg aws.Config) *ssm.Client {
	return ssm.NewFromConfig(cfg)
}

func (aw *AWSconfig) RetrieveSSMValue() (string, string, error) {

	res, err := aw.SSMConfig.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: aws.String(SSM_PATH),
	})

	// NotFound
	if err != nil {
		return "", ECR_NAME, err
	}

	return *res.Parameter.Name, ECR_NAME, nil
}

func (aw *AWSconfig) SetSSMValue() (string, error) {

	_, err := aw.SSMConfig.PutParameter(context.TODO(), &ssm.PutParameterInput{
		Name:  aws.String(SSM_PATH),
		Value: aws.String(ECR_NAME),
		Type:  "String",
	})

	if err != nil {
		return "", err
	}

	return ECR_NAME, nil
}
