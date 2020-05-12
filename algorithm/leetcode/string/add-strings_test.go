package string

import (
	"testing"
)

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{num1: "111", num2: "222"}, want: "333"},
		{name: "test2", args: args{num1: "999", num2: "999"}, want: "1998"},
		{name: "test2", args: args{num1: "408", num2: "5"}, want: "413"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
