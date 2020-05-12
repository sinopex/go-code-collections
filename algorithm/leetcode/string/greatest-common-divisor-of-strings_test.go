package string

import (
	"testing"
)

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{a: 100, b: 72}, want: 4},
		{name: "test2", args: args{a: 64, b: 36}, want: 4},
		{name: "test3", args: args{a: 72, b: 36}, want: 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcdOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{str1: "ABABABAB", str2: "ABAB"}, want: "ABAB"},
		{name: "test2", args: args{str1: "ABABABAB", str2: "ABABAB"}, want: "AB"},
		{name: "test3", args: args{str1: "LEET", str2: "CODE"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
