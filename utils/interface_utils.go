package utils

import (
	"errors"
	"reflect"
)

func InterfaceToArray(actual interface{}) ([]interface{}, error) {
	var res []interface{}
	value := reflect.ValueOf(actual)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return nil, errors.New("parse interface to arrary failed")
	}

	for i := 0; i < value.Len(); i++ {
		res = append(res, value.Index(i).Interface())
	}
	return res, nil
}
