package array

import (
	"testing"
)

func Test_checkStraightLine(t *testing.T) {
	type args struct {
		coordinates [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				coordinates: [][]int{
					{1, 2}, {2, 3}, {3, 4},
					{4, 5}, {5, 6}, {6, 7},
				},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				coordinates: [][]int{
					{0, 1}, {0, 2}, {0, 3},
					{0, 4}, {0, 5}, {0, 6},
				},
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				coordinates: [][]int{
					{1, 0}, {2, 0}, {3, 0},
					{4, 0}, {5, 0}, {6, 0},
				},
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				coordinates: [][]int{
					{-3, 4}, {-2, 3}, {-1, 2}, {0, 1},
				},
			},
			want: true,
		},
		{
			name: "test5",
			args: args{
				coordinates: [][]int{
					{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7},
				},
			},
			want: false,
		},
		{
			name: "test6",
			args: args{
				coordinates: [][]int{
					{0, 0}, {0, 5}, {5, 5}, {5, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStraightLine(tt.args.coordinates); got != tt.want {
				t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
