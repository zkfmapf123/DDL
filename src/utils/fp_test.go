package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type params struct {
	Name string
	Age  string
}

var p = params{
	Name: "leedonggyu",
	Age:  "30",
}

func Test_GetKey(t *testing.T) {

	r1 := GetKey(p, "Name")
	r2 := GetKey(p, "Age")
	r3 := GetKey(p, "personality")

	assert.Equal(t, r1, "leedonggyu")
	assert.Equal(t, r2, "30")
	assert.Equal(t, r3, "")
}

func Test_OKey(t *testing.T) {

	res := OKeys(p)

	assert.Equal(t, len(res), 2)
}

func Test_OValues(t *testing.T) {

	res := OValues(p)

	assert.Equal(t, len(res), 2)
}
