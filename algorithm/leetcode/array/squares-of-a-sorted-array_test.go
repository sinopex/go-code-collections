package array

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{A: []int{-4, -1, 0, 3, 10}}, want: []int{0, 1, 9, 16, 100}},
		{name: "test2", args: args{A: []int{-7, -3, 2, 3, 11}}, want: []int{4, 9, 9, 49, 121}},
		{name: "test3", args: args{A: []int{-2, -2, -2, 2, 2, 2}}, want: []int{4, 4, 4, 4, 4, 4}},
		{name: "test4", args: args{A: []int{0, 2}}, want: []int{0, 4}},
		{name: "test5", args: args{A: []int{0}}, want: []int{0}},
		{name: "test6", args: args{A: []int{-3, -1, 0}}, want: []int{0, 1, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
