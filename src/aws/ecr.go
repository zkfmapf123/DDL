package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

// ////////////////////////////////////////////////////// Params ////////////////////////////////////////////////////////
type S3Params struct {
	Name string `json:"name"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
* Provider : AWS
 */
func GetECRConfig(cfg aws.Config) *ecr.Client {
	return ecr.NewFromConfig(cfg)
}

type S3OptsFuncs func(opts *S3Params)
