package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testConfig = NewAWSCredentials().WithProfile("leedonggyu").MustEnd()

func Test_DefaultAWS(t *testing.T) {

	config := NewAWSCredentials().WithProfile("default").
		WithRegion("us-east-1").
		WithAccessSecretKey("acKey", "seKey")

	assert.Equal(t, config.profile, "default")
	assert.Equal(t, config.region, "us-east-1")
	assert.Equal(t, config.awsAccessKey, "acKey")
	assert.Equal(t, config.awsSecretKey, "seKey")
}

func Test_AWSConfig(t *testing.T) {

	config := NewAWSCredentials().WithProfile("leedonggyu").MustEnd()

	assert.NotEqual(t, config.ssmConfig, nil)
}
