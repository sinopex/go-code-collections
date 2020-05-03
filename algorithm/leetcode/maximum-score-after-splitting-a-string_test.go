package leetcode

import (
	"testing"
)

func Test_maxScore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "test1", s: "011101", want: 5},
		{name: "test2", s: "0111000001", want: 7},
		{name: "test3", s: "011101", want: 5},
		{name: "test4", s: "1111", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.s); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
