package leetcode

import (
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{s: "aba"}, want: true},
		{name: "test2", args: args{s: "abca"}, want: true},
		{name: "test3", args: args{s: "abeca"}, want: false},
		{name: "test4", args: args{s: "11111111"}, want: true},
		{name: "test4", args: args{s: "abcdefgfedcbab"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
