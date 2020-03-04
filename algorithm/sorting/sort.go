package sorting

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func mockUnsortedArray(count, min, max int) (arr []int) {
	for i := 0; i < count; i++ {
		arr = append(arr, min+r.Intn(max-min))
	}
	return
}

type Sorter func(arr []int) []int

func copyArray(arr []int) (sorted []int, length int) {
	if length = len(arr); length > 0 {
		sorted = make([]int, length)
		_ = copy(sorted, arr)
	}

	return
}
