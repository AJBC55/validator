package validator

import "reflect"

func makeStructSlice(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	fields := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, v.Field(i))
	}
	return fields
}

func Validate(s interface{}) bool {
	fields := makeStructSlice(s)
	for _, field := range fields {
		required, ok := field.(RequiredValue)
		if ok {
			if !required.valid() {
				return false
			}
		}
	}
	return true
}
