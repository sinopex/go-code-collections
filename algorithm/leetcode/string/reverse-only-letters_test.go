package string

import (
	"testing"
)

func Test_reverseOnlyLetters(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{S: "ab-cd"}, want: "dc-ba"},
		{name: "test2", args: args{S: "a-bC-dEf-ghIj"}, want: "j-Ih-gfE-dCba"},
		{name: "test3", args: args{S: "Test1ng-Leet=code-Q!"}, want: "Qedo1ct-eeLg=ntse-T!"},
		{name: "test4", args: args{S: "ab-cd-aaaaa-bbbb-eee-1234--aab--ba"}, want: "ab-ba-aeeeb-bbba-aaa-1234--adc--ba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.args.S); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
