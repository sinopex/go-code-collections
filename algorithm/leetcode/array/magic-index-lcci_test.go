package array

import (
	"testing"
)

func Test_findMagicIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{0, 1, 2, 3, 4, 5}}, want: 0},
		{name: "test2", args: args{nums: []int{1, 2, 3, 4, 5}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMagicIndex(tt.args.nums); got != tt.want {
				t.Errorf("findMagicIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
