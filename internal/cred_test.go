package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ddlCreds(t *testing.T) {

	creds, err := GetAWSCredsProfile()

	if err == nil {
		assert.Equal(t, creds.Profile == "", false)
		assert.Equal(t, creds.Region == "", false)
	}
}
