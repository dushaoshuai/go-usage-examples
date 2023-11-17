package funcs

import (
	"reflect"
	"runtime"
	"strings"
)

func GetFuncName(f any) string {
	rv := reflect.ValueOf(f)
	if rv.Kind() != reflect.Func {
		return "not a func"
	}

	pc := uintptr(rv.UnsafePointer())
	fullName := runtime.FuncForPC(pc).Name()
	subStrings := strings.SplitAfter(fullName, ".")
	if l := len(subStrings); l >= 1 {
		return subStrings[l-1]
	}
	return fullName
}
