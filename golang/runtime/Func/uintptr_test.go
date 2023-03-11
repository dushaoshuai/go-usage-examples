package Func_test

import (
	"fmt"
	"reflect"
	"runtime"
)

func f() {}

var f2 = f

func getPc(f any) uintptr {
	return uintptr(reflect.ValueOf(f).UnsafePointer())
}

func Example_uintptr() {
	pc := getPc(f2) // 使用 f2 得到的结果和 f 是一样的
	rf := runtime.FuncForPC(pc)
	fmt.Println(rf.Name())
	fmt.Println(rf.Entry() == pc)
	fmt.Println(rf.FileLine(pc))

	// Output:
	// github.com/dushaoshuai/go-usage-examples/golang/runtime/Func_test.f
	// true
	// /Users/ingtube/Documents/dev/shaouai/go-usage-examples/golang/runtime/Func/uintptr_test.go 9
}
