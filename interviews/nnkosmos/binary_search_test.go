package nnkosmos

import (
	"testing"
)

func Test_binarySearch(t *testing.T) {
	var (
		threeInts = []int{-7, 2, 10}
		fourInts  = []int{7, 8, 20, 45}
	)

	type args struct {
		x      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"noinput", args{nil, 6}, -1},
		{"one_input_found", args{[]int{1}, 1}, 0},
		{"one_input_not_found", args{[]int{1}, 7}, -1},
		{"two_input_found_first", args{[]int{-2, 7}, -2}, 0},
		{"two_input_found_second", args{[]int{-2, 7}, 7}, 1},
		{"two_input_not_found", args{[]int{-2, 7}, 3}, -1},
		{"three_input_found_first", args{threeInts, -7}, 0},
		{"three_input_found_second", args{threeInts, 2}, 1},
		{"three_input_found_third", args{threeInts, 10}, 2},
		{"three_input_not_found", args{threeInts, 101}, -1},
		{"four_input_found_first", args{fourInts, 7}, 0},
		{"four_input_found_second", args{fourInts, 8}, 1},
		{"four_input_found_third", args{fourInts, 20}, 2},
		{"four_input_found_fourth", args{fourInts, 45}, 3},
		{"four_input_not_found", args{fourInts, -10}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.x, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
