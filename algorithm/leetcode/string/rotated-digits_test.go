package string

import (
	"testing"
)

func Test_rotatedDigits(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{N: 10}, want: 4},
		{name: "test2", args: args{N: 2}, want: 1},
		{name: "test3", args: args{N: 100}, want: 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotatedDigits(tt.args.N); got != tt.want {
				t.Errorf("rotatedDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
