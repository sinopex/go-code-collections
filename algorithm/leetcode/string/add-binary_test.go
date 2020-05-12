package string

import (
	"testing"
)

func Test_addBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{name: "test1", a: "11", b: "1", want: "100"},
		{name: "test2", a: "1010", b: "1011", want: "10101"},
		{name: "test3", a: "111", b: "111", want: "1110"},
		{name: "test4", a: "1", b: "111", want: "1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAddBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addBinary("10011111100010101", "1010111000110100011")
	}
}
