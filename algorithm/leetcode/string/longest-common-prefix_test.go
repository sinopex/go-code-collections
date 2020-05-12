package string

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name:"test1",
		// 	args: args{
		// 		strs: []string{"flower","flow","flight"},
		// 	},
		// 	want: "fl",
		// },
		{
			name: "test2",
			args: args{
				strs: []string{"ab", "abc", "abcd"},
			},
			want: "ab",
		},
		{
			name: "test3",
			args: args{
				strs: []string{"ab", "abc", "abcd", ""},
			},
			want: "",
		},
		{
			name: "test4",
			args: args{
				strs: []string{"bb", "abc", "abcd", ""},
			},
			want: "",
		},
		{
			name: "test5",
			args: args{
				strs: []string{"abcd", "abcd", "abcd"},
			},
			want: "abcd",
		},
		{
			name: "test6",
			args: args{
				strs: []string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
