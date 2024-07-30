package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

// ////////////////////////////////////////////////////// Params ////////////////////////////////////////////////////////
type ECRParams struct {
	Name string
	Arn  string
	Uri  string
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
* Provider : AWS
 */
func GetECRConfig(cfg aws.Config) *ecr.Client {
	return ecr.NewFromConfig(cfg)
}

func (aw AWSconfig) RetrievECR(name string) (ECRParams, error) {
	res, err := aw.ECRConfig.DescribeRepositories(context.TODO(), &ecr.DescribeRepositoriesInput{
		RepositoryNames: []string{name},
	})

	if err != nil {
		return ECRParams{}, err
	}

	ecrName := ECRParams{}
	for _, v := range res.Repositories {
		ecrName.Name = *v.RepositoryName
		ecrName.Uri = *v.RepositoryUri
		ecrName.Arn = *v.RepositoryArn
	}

	return ecrName, nil
}

func (aw AWSconfig) CreateECR(name string) (ECRParams, error) {
	fmt.Println(name)

	res, err := aw.ECRConfig.CreateRepository(context.TODO(), &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(name),
	})

	if err != nil {
		return ECRParams{}, err
	}

	return ECRParams{
		Name: *res.Repository.RepositoryName,
		Arn:  *res.Repository.RepositoryArn,
		Uri:  *res.Repository.RepositoryUri,
	}, nil
}
