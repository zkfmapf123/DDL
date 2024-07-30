package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RetrieveSSMValue(t *testing.T) {

	aws, _, _ := New()

	v, _ := aws.RetrieveSSMValue()

	assert.NotEqual(t, v, nil)
}

// func Test_SetSSMValue(t *testing.T) {

// 	aws, _, _ := New()
// 	_, err := aws.SetSSMValuev("test_ecr")

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }
