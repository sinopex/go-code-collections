package array

import (
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
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
			args: args{nums: []int{1, 1, 0, 1, 1, 1}},
			want: 3,
		},
		{
			name: "test2",
			args: args{nums: []int{1, 1, 1, 0, 0, 1, 1, 1, 0, 0}},
			want: 3,
		},
		{
			name: "test3",
			args: args{nums: []int{0, 0, 0}},
			want: 0,
		},
		{
			name: "test4",
			args: args{nums: []int{1, 1, 1, 1, 1, 1, 1, 1}},
			want: 8,
		},
		{
			name: "test5",
			args: args{nums: []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}},
			want: 1,
		},
		{
			name: "test6",
			args: args{[]int{0}},
			want: 0,
		},
		{
			name: "test7",
			args: args{nums: []int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
