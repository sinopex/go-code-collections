package leetcode

import (
	"testing"
)

func Test_mostCommonWord(t *testing.T) {
	type args struct {
		paragraph string
		banned    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
				banned:    []string{"hit"},
			},
			want: "ball",
		},
		{
			name: "test2",
			args: args{
				paragraph: "a ball,a cat,a dog,a big cat,a small dog,a BIG ball",
				banned:    []string{"dog", "cat", "a", "big"},
			},
			want: "ball",
		},
		{
			name: "test3",
			args: args{
				paragraph: "Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. Bob hit a ball, the hit BALL flew far after it was hit. ",
				banned:    []string{"hit"},
			},
			want: "ball",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCommonWord(tt.args.paragraph, tt.args.banned); got != tt.want {
				t.Errorf("mostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
