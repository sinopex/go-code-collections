package tree

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "test",
			input: []int{-10, -3, 0, 5, 9},
			want:  []int{-10, -3, 0, 5, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(sortedArrayToBST(tt.input)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
