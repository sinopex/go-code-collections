package leetcode

import (
	"testing"
)

func Test_countSegments(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{s: "Hello, my name is John"}, want: 5},
		{name: "test2", args: args{s: "a"}, want: 1},
		{name: "test3", args: args{s: " "}, want: 0},
		{name: "test4", args: args{s: "           "}, want: 0},
		{name: "test5", args: args{s: "           abc@123"}, want: 1},
		{name: "test6", args: args{s: " a b c d e f g "}, want: 7},
		{name: "test7", args: args{s: "Of all the gin joints in all the towns in all the world,   "}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments(tt.args.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}
