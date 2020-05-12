package string

import (
	"testing"
)

func isFlipedStringTests() []struct {
	name string
	s1   string
	s2   string
	want bool
} {
	return []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{name: "test1", s1: "waterbottle", s2: "erbottlewat", want: true},
		{name: "test2", s1: "aa", s2: "aba", want: false},
		{name: "test3", s1: "abcd", s2: "bacd", want: false},
		{name: "test4", s1: "abcdefghijklmnopqrstuvwxyz", s2: "opqrstuvwxyzabcdefghijklmn", want: true},
	}
}
func Test_isFlipedString(t *testing.T) {
	for _, tt := range isFlipedStringTests() {
		t.Run(tt.name, func(t *testing.T) {
			isFlipedString(tt.s1, tt.s2)
		})
	}
}

func Test_isFlipedString2(t *testing.T) {
	for _, tt := range isFlipedStringTests() {
		t.Run(tt.name, func(t *testing.T) {
			isFlipedString2(tt.s1, tt.s2)
		})
	}
}

func BenchmarkFlipedString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range isFlipedStringTests() {
			b.Run(tt.name, func(b *testing.B) {
				if got := isFlipedString(tt.s1, tt.s2); got != tt.want {
					b.Errorf("isFlipedString() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func BenchmarkFlipedString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range isFlipedStringTests() {
			b.Run(tt.name, func(b *testing.B) {
				if got := isFlipedString2(tt.s1, tt.s2); got != tt.want {
					b.Errorf("isFlipedString() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
