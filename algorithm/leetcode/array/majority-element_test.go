package array

import (
	"testing"
)

func Test_majorityElement1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement1(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement1() = %v, want %v", got, tt.want)
			}
		})
	}
}
