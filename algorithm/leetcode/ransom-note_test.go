package leetcode

import (
	"testing"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{name: "test1", ransomNote: "a", magazine: "b", want: false},
		{name: "test2", ransomNote: "aa", magazine: "ab", want: false},
		{name: "test3", ransomNote: "aa", magazine: "aab", want: true},
		{name: "test4", ransomNote: "tee", magazine: "teen", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.ransomNote, tt.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
