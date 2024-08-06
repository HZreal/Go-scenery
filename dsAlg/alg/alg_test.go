package alg

/**
 * @Author elastic·H
 * @Date 2024-06-05
 * @File: alg_test.go
 * @Description:
 */

import (
	"testing"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 6, 7, 8, 9, 10}, 7}, 5},
		{"case2", args{[]int{1, 2, 3, 4, 6, 7, 8, 9, 10}, 5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Log("got  ---->  ", got)
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
