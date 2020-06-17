package array

import (
	"testing"
)

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{1, 3, 5, 4, 7}}, want: 3},
		{name: "test2", args: args{nums: []int{2, 2, 2, 2, 2}}, want: 1},
		{name: "test3", args: args{nums: []int{1, 3, 5, 7, 9}}, want: 5},
		{name: "test4", args: args{nums: []int{4, 3, 2, 1}}, want: 1},
		{name: "test5", args: args{nums: []int{1, 2, 3, 0, 4, 5, 6}}, want: 4},
		{name: "test6", args: args{nums: []int{1}}, want: 1},
		{name: "test7", args: args{nums: []int{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
