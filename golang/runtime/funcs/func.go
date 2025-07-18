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
	subStrings := strings.Split(fullName, ".")
	if l := len(subStrings); l >= 1 {
		return subStrings[l-1]
	}
	return fullName
}

// GetCallerName get the name of caller of GetCallerName.
func GetCallerName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return ""
	}

	fullName := fn.Name()
	subStrings := strings.Split(fullName, ".")
	if l := len(subStrings); l >= 1 {
		return subStrings[l-1]
	}
	return fullName
}

// GetCallerNameSkip reports caller name. The argument skip is the number of stack frames
// to ascend, with 0 identifying the caller of GetCallerNameSkip (the same as GetCallerName).
func GetCallerNameSkip(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return ""
	}

	fullName := fn.Name()
	subStrings := strings.Split(fullName, ".")
	if l := len(subStrings); l >= 1 {
		return subStrings[l-1]
	}
	return fullName
}
