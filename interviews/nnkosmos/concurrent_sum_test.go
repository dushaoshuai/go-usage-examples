package nnkosmos

import (
	"testing"
)

func Test_concurrentSum(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test20", args{20, 11}, 210},
		{"test100", args{100, 3}, 5050},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concurrentSum(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("concurrentSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
