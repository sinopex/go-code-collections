package string

import (
	"testing"
)

func Test_canPermutePalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "test1", s: "tactcoa", want: true},
		{name: "test2", s: "abcaa", want: false},
		{name: "test3", s: "abca", want: false},
		{name: "test4", s: "aab", want: true},
		{name: "test5", s: "aabbba", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPermutePalindrome(tt.s); got != tt.want {
				t.Errorf("canPermutePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
