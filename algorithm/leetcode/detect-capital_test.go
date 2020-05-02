package leetcode

import (
	"testing"
)

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{word: "USA"}, want: true},
		{name: "test2", args: args{word: "Google"}, want: true},
		{name: "test3", args: args{word: "leetcode"}, want: true},
		{name: "test4", args: args{word: "LeetCode"}, want: false},
		{name: "test4", args: args{word: "ffffffffffffffffffffF"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.args.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
