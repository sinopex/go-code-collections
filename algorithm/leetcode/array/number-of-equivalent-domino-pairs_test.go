package array

import (
	"testing"
)

func Test_numEquivDominoPairs(t *testing.T) {
	type args struct {
		dominoes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				dominoes: [][]int{
					{1, 2}, {2, 1}, {3, 4}, {5, 6},
				},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				dominoes: [][]int{
					{1, 2}, {1, 2},
				},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				dominoes: [][]int{
					{1, 2}, {1, 3}, {1, 4}, {1, 5},
				},
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				dominoes: [][]int{
					{1, 2}, {1, 2}, {1, 2}, {1, 2},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEquivDominoPairs(tt.args.dominoes); got != tt.want {
				t.Errorf("numEquivDominoPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
