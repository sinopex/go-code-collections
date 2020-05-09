package leetcode

import (
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{haystack: "hello", needle: "ll"}, want: 2},
		{name: "test1", args: args{haystack: "abcdabce", needle: "abce"}, want: 4},
		{name: "test2", args: args{haystack: "aaaaa", needle: "aba"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
