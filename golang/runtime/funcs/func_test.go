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
