package database

import (
	"fmt"
	"reflect"

	"github.com/PatateDu609/matcha/utils"
	"github.com/PatateDu609/matcha/utils/log"
)

func GetColumns[T any](includeOptional bool) []string {
	var t T
	reflectedType := reflect.TypeOf(t)

	if reflectedType.Kind() != reflect.Struct {
		return nil
	}

	num := reflectedType.NumField()
	res := make([]string, 0, num)
	for i := 0; i < num; i++ {
		field := reflectedType.Field(i)

		if !field.IsExported() {
			continue
		}

		if field.Type.Kind() == reflect.Struct && field.Name == "Base" && field.Type.Name() == "Base" && field.Type.PkgPath() == currentPkg {
			res = append(res, GetColumns[Base](includeOptional)...)
			continue
		}

		if !includeOptional && field.Type.Kind() == reflect.Pointer {
			continue
		}

		name, ok := field.Tag.Lookup(dbTag)
		if !ok {
			name = utils.PascalToSnakeCase(field.Name)
			log.Logger.Trace(fmt.Sprintf("no name specified for %s, taking %s instead", field.Name, name))
		}

		res = append(res, name)
	}
	return res
}

func GetInterfaceArray[T any](instanceOf *T) []any {
	reflectedValue := reflect.ValueOf(instanceOf).Elem()
	reflectedType := reflectedValue.Type()
	res := make([]any, 0, reflectedValue.NumField())

	for i := 0; i < reflectedValue.NumField(); i++ {
		field := reflectedValue.Field(i)
		fieldType := reflectedType.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		if field.Kind() == reflect.Struct && fieldType.Name == "Base" && fieldType.Type.Name() == "Base" && fieldType.Type.PkgPath() == currentPkg {
			res = append(res, GetInterfaceArray[Base](field.Addr().Interface().(*Base))...)
			continue
		}

		res = append(res, field.Addr().Interface())
	}

	return res
}

func PrepareValues[T any](data T) []any {
	reflectedValue := reflect.ValueOf(data)

	if reflectedValue.Kind() == reflect.Pointer {
		reflectedValue = reflectedValue.Elem()
	}
	if reflectedValue.Kind() != reflect.Struct {
		log.Logger.Error("Please pass a structure")
		return nil
	}
	reflectedType := reflectedValue.Type()

	res := make([]any, reflectedValue.NumField())

	for i := 0; i < reflectedValue.NumField(); i++ {
		field := reflectedValue.Field(i)
		fieldType := reflectedType.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		res[i] = field.Interface()
	}

	return res
}
