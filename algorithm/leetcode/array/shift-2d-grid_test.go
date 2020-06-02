package array

import (
	"reflect"
	"testing"
)

func Test_shiftGrid(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				grid: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				k: 1,
			},
			want: [][]int{
				{9, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
		},
		{
			name: "test2",
			args: args{
				grid: [][]int{
					{3, 8, 1, 9},
					{19, 7, 2, 5},
					{4, 6, 11, 10},
					{12, 0, 21, 13},
				},
				k: 4,
			},
			want: [][]int{
				{12, 0, 21, 13},
				{3, 8, 1, 9},
				{19, 7, 2, 5},
				{4, 6, 11, 10},
			},
		},
		{
			name: "test3",
			args: args{
				grid: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				k: 9,
			},
			want: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name: "test4",
			args: args{
				grid: [][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}},
				k:    23,
			},
			want: [][]int{{6}, {5}, {1}, {2}, {3}, {4}, {7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftGrid(tt.args.grid, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shiftGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
