package utils

import (
	"reflect"
)

func GetKey[T any](data T, key string) string {

	dataValue := reflect.ValueOf(data)

	field := dataValue.FieldByName(key)
	if field.IsValid() && field.CanInterface() {
		return field.String()
	}

	return ""
}

func OKeys[T any](data T) []string {

	res := reflect.ValueOf(data)
	tpy := res.Type()
	numFields := tpy.NumField()

	keys := []string{}
	for i := 0; i < numFields; i++ {
		keys = append(keys, tpy.Field(i).Name)
	}

	return keys

}

func OValues[T any](data T) []string {

	res := reflect.ValueOf(data)
	tpy := res.Type()
	numFields := tpy.NumField()

	values := []string{}
	for i := 0; i < numFields; i++ {
		values = append(values, res.Field(i).String())
	}

	return values
}
