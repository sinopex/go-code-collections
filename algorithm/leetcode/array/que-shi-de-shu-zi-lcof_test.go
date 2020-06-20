package array

import (
	"testing"
)

func Test_missingNumber4(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{0, 1, 3}}, want: 2},
		{name: "test2", args: args{nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 9}}, want: 8},
		{name: "test23", args: args{nums: []int{0}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber4(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber4() = %v, want %v", got, tt.want)
			}
		})
	}
}
