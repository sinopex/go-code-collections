package array

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				nums:   []int{1, 1, 1, 1, 1, 1, 1},
				target: 1,
			},
			want: 7,
		},
		{
			name: "test4",
			args: args{
				nums:   []int{},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
