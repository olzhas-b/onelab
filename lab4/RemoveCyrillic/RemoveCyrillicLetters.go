package RemoveCyrillic

import (
	"reflect"
)

func rec(structValue reflect.Value) interface{} {
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		for field.Kind() == reflect.Ptr {
			field = field.Elem()
		}
		if field.Kind() == reflect.String {
			field.SetString(withoutCyrillic(field.String()))
		} else if field.Kind() == reflect.Struct {
			rec(field)
		}
	}
	return nil
}
func RemoveCyrillicFromStruct(structInterface interface{}) (err error) {
	val := reflect.ValueOf(structInterface)
	if val.Kind() != reflect.Ptr {
		panic("it is not pointer")
	}
	val = val.Elem()
	if val.Kind() != reflect.Struct {
		panic("points not to struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}
		if field.Kind() == reflect.String {
			// if it is string
			field.SetString(withoutCyrillic(field.String()))
		} else if field.Kind() == reflect.Struct {
			// if it is struct
			rec(field)
		}
	}
	return
}

func withoutCyrillic(s interface{}) (result string) {
	switch s.(type) {
	case string:
		for _, char := range s.(string) {
			if (char < 'а' || char > 'я') && (char < 'А' || char > 'Я') {
				result += string(char)
			}
		}
	}
	return
}
