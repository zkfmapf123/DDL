package aws

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	EXISTS_IMAGE = "hello-world"
	CREATE_IMAGE = "ed-pattern"
)

func Test_Already_Exist_CreateECR(t *testing.T) {

	aws, _, _ := New()

	_, err := aws.CreateECR(EXISTS_IMAGE)
	if err == nil {
		log.Fatalln(err)
	}
}

func Test_RetrieveECR(t *testing.T) {
	aws, _, _ := New()

	ecr, err := aws.RetrievECR(EXISTS_IMAGE)
	if err != nil {
		log.Fatalln(err)
	}

	assert.NotEqual(t, ecr.Arn, nil)
	assert.NotEqual(t, ecr.Name, nil)
	assert.NotEqual(t, ecr.Uri, nil)
}
