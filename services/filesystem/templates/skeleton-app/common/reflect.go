package common

import (
	"reflect"
	"strings"
)

func GetTypeName(v interface{}) string {
	return strings.Replace(reflect.TypeOf(v).String(), "*", "", -1)
}

func GetFieldName(structPoint interface{}, fieldPinter interface{}) (name string) {

	val := reflect.ValueOf(structPoint).Elem()
	val2 := reflect.ValueOf(fieldPinter).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if valueField.Addr().Interface() == val2.Addr().Interface() {
			return val.Type().Field(i).Name
		}
	}

	return
}
