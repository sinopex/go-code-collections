package array

import (
	"reflect"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				mat: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 0},
					{1, 0, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 1},
				},
				k: 3,
			},
			want: []int{2, 0, 3},
		},
		{
			name: "test2",
			args: args{
				mat: [][]int{
					{1, 0, 0, 0},
					{1, 1, 1, 1},
					{1, 0, 0, 0},
					{1, 0, 0, 0},
				},
				k: 2,
			},
			want: []int{0, 2},
		},
		{
			name: "test3",
			args: args{
				mat: [][]int{
					{1, 0},
					{1, 0},
					{1, 0},
					{1, 1},
				},
				k: 4,
			},
			want: []int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kWeakestRows(tt.args.mat, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kWeakestRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
