package array

import (
	"reflect"
	"testing"
)

func Test_masterMind(t *testing.T) {
	type args struct {
		solution string
		guess    string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{solution: "RGBY", guess: "GGRR"}, want: []int{1, 1}},
		{name: "test2", args: args{solution: "RGBY", guess: "RGBY"}, want: []int{4, 0}},
		{name: "test2", args: args{solution: "RRYY", guess: "GBGB"}, want: []int{0, 0}},
		{name: "test2", args: args{solution: "RRYY", guess: "YYRR"}, want: []int{0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := masterMind(tt.args.solution, tt.args.guess); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("masterMind() = %v, want %v", got, tt.want)
			}
		})
	}
}
