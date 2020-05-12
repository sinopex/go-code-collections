package string

import (
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s1: "abc",
				s2: "bca",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s1: "abd",
				s2: "bca",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				s1: "aa",
				s2: "bb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPermutation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CheckPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
