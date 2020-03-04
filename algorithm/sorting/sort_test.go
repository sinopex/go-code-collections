package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases []params
var testFuncs map[string]Sorter

type params struct {
	name string
	args []int
	want []int
}

func init() {
	const length = 1000
	var test6Want = make([]int, length)
	for i := 0; i < length; i++ {
		test6Want[i] = i
	}
	m := make(map[int]struct{}, length)
	for i := 0; i < length; i++ {
		m[i] = struct{}{}
	}
	var test6Args = make([]int, 0)
	for r := range m {
		test6Args = append(test6Args, r)
	}

	testCases = []params{
		{
			name: "test1",
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "test2",
			args: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "test3",
			args: []int{1, 5, 2, 10, 6, 9, 3, 4, 8, 7},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "test4",
			args: []int{105, 115, 285, 564, 981, 902, 503, 444, 208, 736},
			want: []int{105, 115, 208, 285, 444, 503, 564, 736, 902, 981},
		},
		{
			name: "test5",
			args: []int{101, 504, 2005, 100010, 236, 9621, 30012, 4001, 8554, 70},
			want: []int{70, 101, 236, 504, 2005, 4001, 8554, 9621, 30012, 100010},
		},
		{
			name: "test6",
			args: test6Args,
			want: test6Want,
		},
	}

	testFuncs = map[string]Sorter{
		"Bubble":       Bubble,
		"BubbleN":      BubbleN,
		"Quick":        Quick,
		"QuickRefined": QuickRefined,
		"Insert":       Insert,
		"Selection":    Selection,
		"Heap":         Heap,
		"Merge":        Merge,
		"Bucket":       Bucket,
		"Counting":     Counting,
		"Shell":        Shell,
		"Radix":        Radix,
	}
}

func Test_mockUnsortedArray(t *testing.T) {
	type args struct {
		count int
		min   int
		max   int
	}
	tests := []struct {
		name    string
		args    args
		wantArr []int
	}{
		{
			name: "100",
			args: args{
				count: 100,
				min:   0,
				max:   10000,
			},
		},
		{
			name: "1000",
			args: args{
				count: 100,
				min:   0,
				max:   10000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArr := mockUnsortedArray(tt.args.count, tt.args.min, tt.args.max); len(gotArr) != tt.args.count {
				t.Errorf("mockUnsortedArray() = %v, want %v", gotArr, tt.wantArr)
			}
		})
	}
}

func TestSort(t *testing.T) {
	for name, fn := range testFuncs {
		for _, tt := range testCases {
			title := fmt.Sprintf("%s.%s", name, tt.name)
			t.Run(title, func(t *testing.T) {
				if got := fn(tt.args); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("input = %v, %s() = %v, want %v", tt.args, title, got, tt.want)
				}
			})
		}
	}
}
