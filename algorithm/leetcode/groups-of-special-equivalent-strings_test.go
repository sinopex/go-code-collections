package leetcode

import (
	"testing"
)

func Test_numSpecialEquivGroups(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				A: []string{"a", "b", "c", "a", "c", "c"},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				A: []string{"aa", "bb", "ab", "ba"},
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				A: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
			},
			want: 3,
		},
		{
			name: "test4",
			args: args{
				A: []string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSpecialEquivGroups(tt.args.A); got != tt.want {
				t.Errorf("numSpecialEquivGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
