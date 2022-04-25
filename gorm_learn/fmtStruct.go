package gorm_learn

import (
	"fmt"
	"reflect"
)

func FmtStruct(st any) {
	stType := reflect.TypeOf(st)
	if stType == nil {
		return
	}
	if stType.Kind() != reflect.Struct {
		return
	}
	fmt.Println("struct info:")
	fmt.Println()
	fmtName(stType.Name())
	fmt.Println()

	stValue := reflect.ValueOf(st)

	fmt.Println("field info:")
	fmt.Println()
	for i := 0; i < stType.NumField(); i++ {
		field := stType.Field(i)
		value := stValue.Field(i)
		fmtName(field.Name)
		fmtFieldValue(value)
		fmt.Println()
	}
}

func fmtName(name string) {
	fmt.Printf("name: %s\n", name)
}

func fmtFieldValue(value reflect.Value) {
	fmt.Print("value: ")
	switch value.Kind() {
	case reflect.Invalid:
		return
	case reflect.Pointer:
		if value.IsNil() {
			fmt.Println(value.Interface())
		} else {
			fmt.Printf("%+v\n", value.Elem().Interface())
		}
	default:
		fmt.Println(value.Interface())
	}
}
