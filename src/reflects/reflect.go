package reflects

import (
	"api-server/src/api"
	"reflect"
)

func LoadRequest(methodName string) []reflect.Value {
	return InvokeObjectMethod(new(api.Request), methodName)
}

func InvokeObjectMethod(object interface{}, methodName string) []reflect.Value {
	var inputs []reflect.Value
	call := reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
	return call
}
