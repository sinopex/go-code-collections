package linklist

import (
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{head: NewNodeList(1, 1)},
			want: 3,
		},
		{
			name: "test2",
			args: args{head: NewNodeList(0, 0, 0)},
			want: 0,
		},
		{
			name: "test3",
			args: args{head: NewNodeList(1, 0, 0)},
			want: 4,
		},
		{
			name: "test4",
			args: args{head: NewNodeList(1, 1, 1)},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
