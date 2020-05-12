package string

import (
	"testing"
)

func Test_reverseStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{
			name: "test1",
			s:    "abcdefg",
			k:    2,
			want: "bacdfeg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.s, tt.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReverseStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseStr("123abc456def01", 3)
	}
}
