package array

import (
	"reflect"
	"testing"
)

func Test_sumZero(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{n: 5}, want: []int{1, 2, 0, -2, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumZero(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
