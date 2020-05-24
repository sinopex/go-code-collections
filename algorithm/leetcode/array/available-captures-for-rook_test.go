package array

import (
	"testing"
)

func Test_numRookCaptures(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				board: [][]byte{
					{'.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', 'p', '.', '.', '.', '.'},
					{'.', '.', '.', 'R', '.', '.', '.', 'p'},
					{'.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', 'p', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				board: [][]byte{
					{'.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
					{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
					{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
					{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
					{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRookCaptures(tt.args.board); got != tt.want {
				t.Errorf("numRookCaptures() = %v, want %v", got, tt.want)
			}
		})
	}
}
