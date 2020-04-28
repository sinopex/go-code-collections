package leetcode

import (
	"testing"
)

func Test_numberOfSteps(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{name: "test1", num: 14, want: 6},
		{name: "test2", num: 8, want: 4},
		{name: "test3", num: 377, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps(tt.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
