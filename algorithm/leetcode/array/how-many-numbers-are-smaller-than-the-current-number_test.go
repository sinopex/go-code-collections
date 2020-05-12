package array

import (
	"reflect"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{nums: []int{8, 1, 2, 2, 3}}, want: []int{4, 0, 1, 1, 3}},
		{name: "test2", args: args{nums: []int{6, 5, 4, 8}}, want: []int{2, 1, 0, 3}},
		{name: "test3", args: args{nums: []int{7, 7, 7, 7}}, want: []int{0, 0, 0, 0}},
		{name: "test4", args: args{nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
