package internal

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const SSMKEY = "/ddl/info/secrets"

type LambdaECRMap struct {
	Arn              string `json:"arn"`
	ECRName          string `json:"ecrName"`
	UpdateAuthor     string `json:"update_author"`
	TotalDeployCount string `json:"total_deploy"`
}

type SSMValueParams struct {
	LambdaECRMap map[string]LambdaECRMap `json:"lambda_ecr_map"`
	CreatedAt    string                  `json:"created_at"`
	UpdatedAt    string                  `json:"updated_at"`
}

func (p *AWSClientParams) InitSSM() (string, error) {

	ssmParams := SSMValueParams{}

	_, err := p.ssmConfig.PutParameter(context.TODO(), &ssm.PutParameterInput{
		Name:  aws.String(SSMKEY),
		Value: aws.String(MustJsonToString(ssmParams)),
		Type:  types.ParameterTypeSecureString,
	})

	if err != nil {
		return SSMKEY, err
	}

	return SSMKEY, err
}

func (p *AWSClientParams) GetSSMKeyValue() (string, string, error) {

	res, err := p.ssmConfig.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(SSMKEY),
		WithDecryption: aws.Bool(true),
	})

	// Parameter Store에서 값을 가져오는데 실패한 경우
	if err != nil {
		return SSMKEY, "", err
	}

	return SSMKEY, *res.Parameter.Value, err
}

// 람다 함수 생성할때
// 람다 함수 배포할때
// 람다 함수 삭제할때
func (p *AWSClientParams) UpdateSSMValue(lambdaName string, lambdaInfo LambdaECRMap) {

}

func (p *AWSClientParams) IsMatchNotExistsKey(err error) bool {
	return strings.Contains(err.Error(), ErrSSMParameterNotFound)
}
