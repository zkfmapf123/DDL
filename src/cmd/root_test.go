package cmd

import (
	"log"
	"testing"
)

func Test_setAWSProfile(t *testing.T) {
	err := setAWSProfile()
	if err != nil {
		log.Fatalln(err)
	}
}
