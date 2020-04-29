package leetcode

import (
	"testing"
)

func Test_sortString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "test", s: "aaaabbbbcccc", want: "abccbaabccba"},
		{name: "test", s: "rat", want: "art"},
		{name: "test", s: "gggggggg", want: "gggggggg"},
		{name: "test", s: "leetcode", want: "cdelotee"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
