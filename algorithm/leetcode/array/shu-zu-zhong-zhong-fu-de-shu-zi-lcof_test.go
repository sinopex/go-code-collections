package array

import (
	"testing"
)

func Test_findRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5, 3},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				nums: []int{2, 3, 1, 0, 3, 5, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
