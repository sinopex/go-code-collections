package array

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{S: "abbxxxxzyy"},
			want: [][]int{{3, 6}},
		},
		{
			name: "test2",
			args: args{S: "abcdddeeeeaabbbcd"},
			want: [][]int{{3, 5}, {6, 9}, {12, 14}},
		},
		{
			name: "test3",
			args: args{S: "abc"},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largeGroupPositions(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}
