package array

import (
	"testing"
)

func Test_tictactoe(t *testing.T) {
	type args struct {
		moves [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				moves: [][]int{
					{0, 0}, {2, 0},
					{1, 1}, {2, 1},
					{2, 2},
				},
			},
			want: "A",
		},
		{
			name: "test2",
			args: args{
				moves: [][]int{
					{0, 0}, {1, 1},
					{0, 1}, {0, 2},
					{1, 0}, {2, 0},
				},
			},
			want: "B",
		},
		{
			name: "test3",
			args: args{
				moves: [][]int{
					{0, 0}, {1, 1},
					{2, 0}, {1, 0},
					{1, 2}, {2, 1},
					{0, 1}, {0, 2},
					{2, 2},
				},
			},
			want: "Draw",
		},
		{
			name: "test4",
			args: args{
				moves: [][]int{
					{0, 0}, {1, 1},
				},
			},
			want: "Pending",
		},
		{
			name: "test5",
			args: args{
				moves: [][]int{
					{1, 1}, {2, 0},
					{0, 2}, {0, 1},
					{0, 0}, {2, 2},
					{2, 1}, {1, 2},
				},
			},
			want: "Pending",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tictactoe(tt.args.moves); got != tt.want {
				t.Errorf("tictactoe() = %v, want %v", got, tt.want)
			}
		})
	}
}
