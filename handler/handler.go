package handler

import (
	"fmt"
	"reflect"
	"time"
)

func MakeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}

func ScanResult(length int, currentVal []interface{}, columns []string) map[string]interface{} {
	value := make(map[string]interface{})
	for i := 0; i < length; i++ {
		key := columns[i]
		val := *(currentVal[i]).(*interface{})
		if val == nil {
			value[key] = nil
			continue
		}
		vType := reflect.TypeOf(val)
		switch vType.String() {
		case "int64":
			value[key] = val.(int64)
		case "string":
			value[key] = val.(string)
		case "time.Time":
			value[key] = val.(time.Time)
		case "[]uint8":
			value[key] = string(val.([]uint8))
		default:
			fmt.Printf("unsupport data type '%s' now\n", vType)
			// TODO remember add other data type
		}
	}
	return value
}
