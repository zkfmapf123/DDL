package internal

import (
	"fmt"
	"testing"
)

func Test_WrapError(t *testing.T) {

	for i := 0; i < 10; i++ {
		WithError(fmt.Sprintf("hello world - %d", i))

	}

	// TraceError()
}
