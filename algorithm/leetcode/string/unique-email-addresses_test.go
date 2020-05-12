package string

import (
	"testing"
)

func Test_numUniqueEmails(t *testing.T) {
	type args struct {
		emails []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				emails: []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				emails: []string{"123@qq.com", "1.2.3@qq.com", "1..2..3+wong@qq.com"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numUniqueEmails(tt.args.emails); got != tt.want {
				t.Errorf("numUniqueEmails() = %v, want %v", got, tt.want)
			}
		})
	}
}
