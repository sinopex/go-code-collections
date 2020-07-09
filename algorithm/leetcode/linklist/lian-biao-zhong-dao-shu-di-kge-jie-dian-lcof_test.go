package linklist

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				head: NewNodeList(1, 2, 3, 4, 5),
				k:    2,
			},
			want: NewNodeList(4, 5),
		},
		{
			name: "test1",
			args: args{
				head: NewNodeList(3, 4, 5),
				k:    3,
			},
			want: NewNodeList(3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
