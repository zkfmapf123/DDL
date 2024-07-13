package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type LambdaParams struct {
	Name string            `json:"name"`
	Tags map[string]string `json:"tags"`
}

type LambdaOutputParams struct {
	Name         string            `json:"name"`
	Arn          string            `json:"arn"`
	Architecture string            `json:"architecture"`
	CodeSize     string            `json:"codeSize"`
	Desc         string            `json:"desc"`
	ImageUrl     string            `json:"imageUrl"`
	Tags         map[string]string `json:"tags"`
}

func (aw *AWSconfig) CreateLambda() {

}

func retriveLambda(l lambda.Client, parmas LambdaParams) (LambdaOutputParams, error) {
	res, err := l.GetFunction(context.TODO(), &lambda.GetFunctionInput{
		FunctionName: aws.String(parmas.Name),
	})

	if err != nil {
		return LambdaOutputParams{}, nil
	}

	fmt.Println(res.Configuration.Architectures)

	// res.Configuration.
	return LambdaOutputParams{
		Name: *res.Configuration.FunctionName,
		Arn:  *res.Configuration.FunctionArn,
	}, nil

}
