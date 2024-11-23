package internal

import (
	"fmt"
	"testing"
)

func Test_GetLambda(t *testing.T) {

	v, err := testConfig.GetLambda()

	if err != nil {
		panic(err)
	}

	fmt.Println(v)

	// b := testConfig.IsMatchNotExistsKey(err)
	// assert.Equal(t, b, true)

	// if b {
	// 	t.Log("key not found")
	// 	err := testConfig.InitSSM()

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

}
