package array

import (
	"reflect"
	"testing"
)

func Test_createTargetArray(t *testing.T) {
	type args struct {
		nums  []int
		index []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums:  []int{0, 1, 2, 3, 4},
				index: []int{0, 1, 2, 2, 1},
			},
			want: []int{0, 4, 1, 3, 2},
		},
		{
			name: "test2",
			args: args{
				nums:  []int{1, 2, 3, 4, 0},
				index: []int{0, 1, 2, 3, 0},
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTargetArray2(tt.args.nums, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTargetArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCreateTargetArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createTargetArray([]int{1, 2, 3, 4, 0, 1, 2, 3, 4, 5}, []int{0, 3, 4, 5, 6, 7, 1, 8, 2, 3})
	}
}
