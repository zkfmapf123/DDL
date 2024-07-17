package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/zkfmapf123/DDL/src/utils"
)

var (
	ATTACH_TAGS = map[string]string{
		"DDL":         "true",
		"CreatedAt":   utils.GetCurrentTime(),
		"UpdatedAt":   utils.GetCurrentTime(),
		"UpdateCount": "0",
	}
)

type LambdaParams struct {
	Name        string
	Achitecture string
	Tags        map[string]string
	Runtime     string // latest
	Mem         string
	Storage     string
	TimeLimit   int
	Policy      string

	// Option (TOBE)
	// network 	struct
	// gateway  struct
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

func (aw *AWSconfig) CreateLambda(params LambdaParams) {

}

func (aw *AWSconfig) DeleteLambda(params LambdaParams) {

}

func retriveLambda(l lambda.Client, parmas LambdaParams) (LambdaOutputParams, error) {
	res, err := l.GetFunction(context.TODO(), &lambda.GetFunctionInput{
		FunctionName: aws.String(parmas.Name),
	})

	if err != nil {
		return LambdaOutputParams{}, err
	}

	tagRes, err := l.ListTags(context.TODO(), &lambda.ListTagsInput{
		Resource: res.Configuration.FunctionArn,
	})

	if err != nil {
		return LambdaOutputParams{}, nil
	}

	return getLambdaFunc(*res.Configuration, *tagRes), nil
}

func retriveLambdas(l lambda.Client) ([]LambdaOutputParams, error) {

	var marker *string
	lambdaList := []LambdaOutputParams{}

	for {
		res, err := l.ListFunctions(context.TODO(), &lambda.ListFunctionsInput{Marker: marker})

		if err != nil {
			return []LambdaOutputParams{}, err
		}

		for _, fn := range res.Functions {

			tagRes, err := l.ListTags(context.TODO(), &lambda.ListTagsInput{
				Resource: fn.FunctionArn,
			})

			if err != nil {
				return []LambdaOutputParams{}, err
			}

			lambdaList = append(lambdaList, getLambdaFunc(fn, *tagRes))
		}

		if res.NextMarker != nil {
			marker = res.NextMarker
		} else {
			break
		}
	}

	return lambdaList, nil
}

func getLambdaFunc(attr types.FunctionConfiguration, tags lambda.ListTagsOutput) LambdaOutputParams {

	var tagMap = make(map[string]string, len(tags.Tags))

	for k, v := range tags.Tags {
		tagMap[k] = v
	}

	return LambdaOutputParams{
		Name: *attr.FunctionName,
		Arn:  *attr.FunctionArn,
		Tags: tagMap,
	}
}
