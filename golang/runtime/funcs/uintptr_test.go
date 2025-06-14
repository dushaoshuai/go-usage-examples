package funcs

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
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

	file, line := rf.FileLine(pc)
	fileSplits := strings.Split(file, "/")
	file = `/` + filepath.Join(fileSplits[len(fileSplits)-5:]...)
	fmt.Println(file, line)

	// Output:
	// github.com/dushaoshuai/go-usage-examples/golang/runtime/funcs.f
	// true
	// /go-usage-examples/golang/runtime/funcs/uintptr_test.go 11
}
