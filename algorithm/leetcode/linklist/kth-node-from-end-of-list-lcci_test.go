package linklist

import (
	"testing"
)

func Test_kthToLast(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				head: NewNodeList(1, 2, 3, 4, 5),
				k:    2,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				head: NewNodeList(1, 5),
				k:    2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthToLast(tt.args.head, tt.args.k); got != tt.want {
				t.Errorf("kthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
