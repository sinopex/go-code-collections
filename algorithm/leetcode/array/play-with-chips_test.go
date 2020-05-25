package array

import (
	"testing"
)

func Test_minCostToMoveChips(t *testing.T) {
	type args struct {
		chips []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{chips: []int{1, 2, 3}}, want: 1},
		{name: "test2", args: args{chips: []int{2, 2, 2, 3, 3}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostToMoveChips(tt.args.chips); got != tt.want {
				t.Errorf("minCostToMoveChips() = %v, want %v", got, tt.want)
			}
		})
	}
}
