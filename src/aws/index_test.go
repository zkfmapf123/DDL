package aws

import (
	"log"
	"testing"
)

func Test_NewAWS(t *testing.T) {

	_, _, err := New()

	if err != nil {
		log.Fatalln(err)
	}
}
