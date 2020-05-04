package leetcode

import (
	"testing"
)

func Test_checkRecord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "test1", s: "PPALLP", want: true},
		{name: "test2", s: "PPALLL", want: false},
		{name: "test2", s: "LALL", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRecord(tt.s); got != tt.want {
				t.Errorf("checkRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
