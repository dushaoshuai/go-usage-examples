package nnkosmos

import (
	"reflect"
	"testing"

	"github.com/dushaoshuai/go-usage-examples/golang/runtime/funcs"
	"golang.org/x/exp/slices"
)

func Test_intersectSolution(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil1", args{nil, []int{1, 2}}, nil},
		{"nil2", args{[]int{1, 2, 3, 1}, nil}, nil},
		{"bothNil", args{nil, nil}, nil},
		{"intersect", args{[]int{-2, 0, 1}, []int{-2, 3, 2}}, []int{-2}},
		{"intersectDuplication", args{[]int{1, -2, 0, 10, 100, -7}, []int{0, -100, 20, 3, 4, 7, -7}}, []int{0, -7}},
	}
	for _, tt := range tests {
		slices.Sort(tt.want)
	}

	solutions := []intersectSolution{
		intersectBruteForce,
		intersectMap,
		intersectSortAndTwoPointer,
		intersectSortAndBinarySearch,
	}
	for _, f := range solutions {
		fName := funcs.GetFuncName(f)
		for _, tt := range tests {
			t.Run(fName+":"+tt.name, func(t *testing.T) {
				got := f(tt.args.a, tt.args.b)
				slices.Sort(got)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s() = %v, want %v", fName, got, tt.want)
				}
			})
		}
	}
}
