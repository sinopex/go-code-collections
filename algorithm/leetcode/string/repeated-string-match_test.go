package string

import (
	"testing"
)

func Test_repeatedStringMatch(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{A: "abcd", B: "cdabcdab"}, want: 3},
		{name: "test2", args: args{A: "abcd", B: "cdabcab"}, want: -1},
		{name: "test3", args: args{A: "abcd", B: "abcd"}, want: 1},
		{name: "test3", args: args{A: "abc", B: "cabcabca"}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("repeatedStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
