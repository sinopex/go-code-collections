package string

import (
	"testing"
)

func Test_freqAlphabets(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "test1", s: "10#11#12", want: "jkab"},
		{name: "test2", s: "1326#", want: "acz"},
		{name: "test3", s: "25#", want: "y"},
		{name: "test4", s: "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#", want: "abcdefghijklmnopqrstuvwxyz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := freqAlphabets(tt.s); got != tt.want {
				t.Errorf("freqAlphabets() = %v, want %v", got, tt.want)
			}
		})
	}
}
