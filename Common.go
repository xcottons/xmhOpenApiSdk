package xmhOpenApiSdk

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ToStr(param interface{}) string {
	if param == nil {
		return ""
	}
	switch v := param.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		of := reflect.ValueOf(param)
		for (of.Kind() == reflect.Ptr && !of.IsNil()) || (of.Kind() == reflect.Interface && !of.IsNil()) {
			of = of.Elem()
		}
		if of.Kind() == reflect.Slice || of.Kind() == reflect.Map || of.Kind() == reflect.Struct {
			bytes, err := json.Marshal(param)
			if err != nil {
				return fmt.Sprintf("error: %v", err)
			}
			return string(bytes)
		}
		return fmt.Sprintf("%v", param)
	}
}
