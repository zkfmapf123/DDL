package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

const (
	SSM_PATH = "/ddl/store/ecr"
)

/*
* Provider : AWS
 */
func GetSSMConfig(cfg aws.Config) *ssm.Client {
	return ssm.NewFromConfig(cfg)
}

func (aw *AWSconfig) RetrieveSSMValue() (string, error) {

	res, err := aw.SSMConfig.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: aws.String(SSM_PATH),
	})

	// NotFound
	if err != nil {
		return "", err
	}

	return *res.Parameter.Name, nil
}

func (aw *AWSconfig) SetSSMValuev(ecrName string) (string, error) {

	_, err := aw.SSMConfig.PutParameter(context.TODO(), &ssm.PutParameterInput{
		Name:  aws.String(SSM_PATH),
		Value: &ecrName,
		Type:  "String",
	})

	if err != nil {
		return "", err
	}

	return ecrName, nil
}
