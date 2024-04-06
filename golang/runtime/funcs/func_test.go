package funcs

import (
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
