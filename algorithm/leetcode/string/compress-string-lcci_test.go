package string

import (
	"testing"
)

func Test_compressString(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{name: "test1", S: "abbccd", want: "abbccd"},
		{name: "test1", S: "aabcccccaaa", want: "a2b1c5a3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.S); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompressString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compressString("aaaabbedbbbhjkkkvvvvvvvvvvvhhhhhhuuukajdfnahq")
	}
}
