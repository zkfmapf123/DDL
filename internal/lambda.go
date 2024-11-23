package internal

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type LambdaParameter struct {
	Architecture []string
	FunctionName string
	CodeSize     string
	Desc         string
	Env          string
	Arn          string
	LastModified string
	MemorySize   string
	Runtime      string
	Timeout      string
}

func (p *AWSClientParams) GetLambda() (LambdaParameter, error) {

	functionList, err := p.lambdaConfig.ListFunctions(context.TODO(), &lambda.ListFunctionsInput{})
	if err != nil {
		return LambdaParameter{}, err
	}

	fnList := make([]LambdaParameter, len(functionList.Functions))
	for i, fn := range functionList.Functions {

		for _, arch := fn.Architectures {
			
		}
		arch := fn.Architectures

		fnList[i] = LambdaParameter{
			Architecture: ,
		}

		fmt.Println(*fn.FunctionName)

	}
	return LambdaParameter{}, err
}

// func filterLambda() {

// }
