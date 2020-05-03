package leetcode

import (
	"testing"
)

func Test_countBinarySubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "test1", s: "00110011", want: 6},
		{name: "test1", s: "00001111000111111", want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBinarySubstrings(tt.s); got != tt.want {
				t.Errorf("countBinarySubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
