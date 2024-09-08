package internal

import "encoding/json"

func MustJsonToString(v interface{}) string {

	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func MusStringToJson[T any](v string) T {

	var res T
	err := json.Unmarshal([]byte(v), &res)
	if err != nil {
		panic(err)
	}

	return res
}
