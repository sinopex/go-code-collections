package array

import (
	"testing"
)

func Test_distanceBetweenBusStops(t *testing.T) {
	type args struct {
		distance    []int
		start       int
		destination int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				distance:    []int{1, 2, 3, 4},
				start:       0,
				destination: 2,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				distance:    []int{1, 2, 3, 4},
				start:       0,
				destination: 3,
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				distance:    []int{7, 10, 1, 12, 11, 14, 5, 0},
				start:       7,
				destination: 2,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceBetweenBusStops(tt.args.distance, tt.args.start, tt.args.destination); got != tt.want {
				t.Errorf("distanceBetweenBusStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
