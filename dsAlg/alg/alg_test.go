package alg

/**
 * @Author elastic·H
 * @Date 2024-06-05
 * @File: alg_test.go
 * @Description:
 */

import (
	"reflect"
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

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{3, 2, 2, 3}, 3}, 2},
		{"case2", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})

		// t.Run(tt.name, func(t *testing.T) {
		// 	if got := removeElement2(tt.args.nums, tt.args.val); got != tt.want {
		// 		t.Errorf("removeElement() = %v, want %v", got, tt.want)
		// 	}
		// })
	}
}

func Test_sortedSquares(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{-4, -1, 0, 3, 10}}, []int{0, 1, 9, 16, 100}},
		{"case2", args{[]int{-7, -3, 2, 3, 11}}, []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
