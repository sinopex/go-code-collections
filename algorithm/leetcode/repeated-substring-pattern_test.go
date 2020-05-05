package leetcode

import (
	"testing"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "test1", s: "abcabc", want: true},
		{name: "test2", s: "abccba", want: false},
		{name: "test3", s: "abcabcabcabc", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
