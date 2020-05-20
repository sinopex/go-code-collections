package array

import (
	"testing"
)

func Test_isUnique(t *testing.T) {
	type args struct {
		astr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{astr: "leetcode"}, want: false},
		{name: "test2", args: args{astr: "abc"}, want: true},
		{name: "test3", args: args{astr: "www"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique(tt.args.astr); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
