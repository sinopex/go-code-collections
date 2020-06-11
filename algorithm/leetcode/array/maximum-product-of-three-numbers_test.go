package array

import (
	"testing"
)

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{1, 2, 3}}, want: 6},
		{name: "test2", args: args{nums: []int{1, 2, 3, 4}}, want: 24},
		{name: "test3", args: args{nums: []int{18, 1, 5, 17, 18}}, want: 5508},
		{name: "test4", args: args{nums: []int{-18, 1, 5, 17, 18}}, want: 1530},
		{name: "test5", args: args{nums: []int{-6, 7, 2, -4, 5}}, want: 168},
		{name: "test6", args: args{nums: []int{0, 1, 9}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
