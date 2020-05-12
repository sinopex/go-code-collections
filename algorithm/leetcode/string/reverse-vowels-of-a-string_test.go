package string

import (
	"testing"
)

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "test1", s: "hello", want: "holle"},
		{name: "test2", s: "ooooeeee", want: "eeeeoooo"},
		{name: "test3", s: "abce", want: "ebca"},
		{name: "test4", s: "dbce", want: "dbce"},
		{name: "test5", s: "aA", want: "Aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
