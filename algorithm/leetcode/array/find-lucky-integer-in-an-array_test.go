package array

import (
	"testing"
)

func Test_findLucky(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{arr: []int{2, 2, 3, 4}}, want: 2},
		{name: "test2", args: args{arr: []int{1, 2, 2, 3, 3, 3}}, want: 3},
		{name: "test3", args: args{arr: []int{5}}, want: -1},
		{name: "test4", args: args{arr: []int{7, 7, 7, 7, 7, 7, 7}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLucky(tt.args.arr); got != tt.want {
				t.Errorf("findLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
