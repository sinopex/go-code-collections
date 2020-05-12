package string

import (
	"testing"
)

func Test_subtractProductAndSum(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "test1", n: 234, want: 15},
		{name: "test2", n: 4421, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractProductAndSum(tt.n); got != tt.want {
				t.Errorf("subtractProductAndSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
