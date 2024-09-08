package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type jsonTest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

func Test_JsonUtils(t *testing.T) {

	jt := jsonTest{
		Name: "leedonggyu",
		Age:  31,
		Job:  "devops",
	}

	res := MustJsonToString(jt)

	assert.NotEqual(t, res, nil)

	res2 := MusStringToJson[jsonTest](res)

	assert.Equal(t, res2.Name, "leedonggyu")
	assert.Equal(t, res2.Age, 31)
	assert.Equal(t, res2.Job, "devops")
}
