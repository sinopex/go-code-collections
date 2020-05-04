package leetcode

import (
	"testing"
)

func Test_reformat(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "test1", s: "123abc", want: "1a2b3c"},
		{name: "test2", s: "abc", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformat(tt.s); got != tt.want {
				t.Errorf("reformat() = %v, want %v", got, tt.want)
			}
		})
	}
}
