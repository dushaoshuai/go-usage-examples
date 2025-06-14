package funcs

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

func TestGetFuncName(t *testing.T) {
	type args struct {
		f any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{"hello"}, "not a func"},
		{"int", args{1}, "not a func"},
		{"f", args{func() {}}, "func1"},
		{"f", args{func() {}}, "func2"},
		{"time.Now", args{time.Now}, "Now"},
		{"(*singleflight.Group).Do", args{new(singleflight.Group).Do}, "Do-fm"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFuncName(tt.args.f); got != tt.want {
				t.Errorf("GetFuncName() = %v, want %v", got, tt.want)
			}
		})
	}
}

type structMethods struct{}

func (structMethods) Method1() {}
func (structMethods) Method2() {}

func Example_get_methods() {
	var sm structMethods

	rv1 := reflect.ValueOf(sm.Method1)
	pc1 := runtime.FuncForPC(uintptr(rv1.UnsafePointer()))
	fmt.Println(pc1.Name())

	rv2 := reflect.ValueOf(sm.Method2)
	pc2 := runtime.FuncForPC(uintptr(rv2.UnsafePointer()))
	fmt.Println(pc2.Name())

	// Output:
	// github.com/dushaoshuai/go-usage-examples/golang/runtime/funcs.structMethods.Method1-fm
	// github.com/dushaoshuai/go-usage-examples/golang/runtime/funcs.structMethods.Method2-fm
}

func testFunc1() string {
	return GetCallerName()
}

func testFunc2() string {
	return GetCallerName()
}

func TestGetCallerName(t *testing.T) {
	testFunc3 := func() string {
		return GetCallerName()
	}

	tests := []struct {
		name string
		fn   func() string
		want string
	}{
		{"f1", testFunc1, "testFunc1"},
		{"f2", testFunc2, "testFunc2"},
		{"f3", testFunc3, "func1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(); got != tt.want {
				t.Errorf("GetCallerName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testF20(skip int) string { return testF21(skip) }
func testF21(skip int) string { return testF22(skip) }
func testF22(skip int) string { return testF23(skip) }
func testF23(skip int) string { return GetCallerNameSkip(skip) }

func TestGetCallerNameSkip(t *testing.T) {
	tests := []struct {
		name string
		skip int
		want string
	}{
		{"skip 0", 0, "testF23"},
		{"skip 1", 1, "testF22"},
		{"skip 2", 2, "testF21"},
		{"skip 3", 3, "testF20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testF20(tt.skip); got != tt.want {
				t.Errorf("GetCallerNameSkip() = %v, want %v", got, tt.want)
			}
		})
	}
}
