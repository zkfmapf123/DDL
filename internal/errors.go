package internal

import (
	"fmt"

	"github.com/pkg/errors"
)

// //////////////////////////// Define Errors //////////////////////////////
var (
	ErrEmptyProfile = "profile is empty"
	ErrEmptyRegion  = "region is empty"
	ErrEmptyKey     = "access key or secret key is empty"

	ErrSSMParameterNotFound = "ParameterNotFound"
)

var ErrorFamily = errors.New("start ErrorStackTrace")

func WithError(msg string) {
	ErrorFamily = errors.Wrap(ErrorFamily, msg)
}

func TraceError() {
	fmt.Printf("%+v\n", ErrorFamily)
}

func TracePanic() {
	panic(ErrorFamily)
}
