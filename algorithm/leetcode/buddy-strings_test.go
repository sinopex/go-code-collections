package leetcode

import (
	"testing"
)

func Test_buddyStrings(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{A: "ab", B: "ba"}, want: true},
		{name: "test2", args: args{A: "ab", B: "ab"}, want: false},
		{name: "test3", args: args{A: "aaaaaaabc", B: "aaaaaaacb"}, want: true},
		{name: "test4", args: args{A: "", B: "aa"}, want: false},
		{name: "test5", args: args{A: "aa", B: "aa"}, want: true},
		{name: "test5", args: args{A: "aab", B: "aac"}, want: false},
		{name: "test5", args: args{A: "abc", B: "acb"}, want: true},
		{name: "test5", args: args{A: "abab", B: "abab"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
