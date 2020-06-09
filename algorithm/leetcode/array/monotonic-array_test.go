package array

import (
	"testing"
)

func Test_isMonotonic(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{A: []int{1, 2, 2, 3}},
			want: true,
		},
		{
			name: "test2",
			args: args{A: []int{6, 5, 4, 4}},
			want: true,
		},
		{
			name: "test3",
			args: args{A: []int{1, 2, 4, 5}},
			want: true,
		},
		{
			name: "test4",
			args: args{A: []int{1, 1, 1, 1, 1}},
			want: true,
		},
		{
			name: "test5",
			args: args{A: []int{1, 3, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.A); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
