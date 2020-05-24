package array

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				A: []int{3, 1, 2, 4},
			},
			want: []int{2, 4, 3, 1},
		},
		{
			name: "test2",
			args: args{
				A: []int{3, 3, 3, 4, 4, 4},
			},
			want: []int{4, 4, 4, 3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
