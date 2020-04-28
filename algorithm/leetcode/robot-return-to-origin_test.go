package leetcode

import (
	"testing"
)

func Test_judgeCircle(t *testing.T) {
	tests := []struct {
		name  string
		moves string
		want  bool
	}{
		{name: "test1", moves: "UD", want: true},
		{name: "test2", moves: "LL", want: false},
		{name: "test3", moves: "LLLUU", want: false},
		{name: "test3", moves: "ULDR", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeCircle(tt.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
