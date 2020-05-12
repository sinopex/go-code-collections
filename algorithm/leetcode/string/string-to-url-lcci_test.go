package string

import (
	"testing"
)

func Test_replaceSpaces(t *testing.T) {
	type args struct {
		S      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{S: "Mr John Smith    ", length: 13}, want: "Mr%20John%20Smith"},
		{name: "test2", args: args{S: "               ", length: 5}, want: "%20%20%20%20%20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpaces(tt.args.S, tt.args.length); got != tt.want {
				t.Errorf("replaceSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
