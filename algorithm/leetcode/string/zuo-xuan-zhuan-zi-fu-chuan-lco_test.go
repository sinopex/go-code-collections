package string

import (
	"testing"
)

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{s: "", n: 2}, want: ""},
		{name: "test1", args: args{s: "abcdefg", n: 2}, want: "cdefgab"},
		{name: "test2", args: args{s: "abcdefg", n: 9}, want: "cdefgab"},
		{name: "test3", args: args{s: "lrloseumgh", n: 6}, want: "umghlrlose"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
