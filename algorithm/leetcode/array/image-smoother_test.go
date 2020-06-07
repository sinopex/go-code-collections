package array

import (
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	type args struct {
		M [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				M: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "test2",
			args: args{
				M: [][]int{
					{4, 4, 5},
					{1, 1, 1},
					{5, 5, 5},
				},
			},
			want: [][]int{
				{2, 2, 2},
				{3, 3, 3},
				{3, 3, 3},
			},
		},
		{
			name: "test3",
			args: args{
				M: [][]int{
					{2, 3, 4},
					{5, 6, 7},
					{8, 9, 10},
					{11, 12, 13},
					{14, 15, 16},
				},
			},
			want: [][]int{
				{4, 4, 5},
				{5, 6, 6},
				{8, 9, 9},
				{11, 12, 12},
				{13, 13, 14},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageSmoother(tt.args.M); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
