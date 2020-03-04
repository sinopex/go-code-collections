package heap

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sinopex/go-code-collections/pkg/assert"
)

// 堆排序
func TestSorting(t *testing.T) {
	// 待排序的slice
	dataset := []int{19, 12, 18, 4, 2, 3, 1, 5, 11, 23, 28, 23, 14, 22, 9, 15}
	h := new(Sorting)
	// 全都入堆
	for _, v := range dataset {
		h.Push(v)
	}
	sort.Sort(sort.IntSlice(dataset))

	for _, v := range dataset {
		n, err := h.Pop()
		assert.Nil(t, err)
		assert.Equal(t, v, n)
	}
	fmt.Println()
	assert.Equal(t, 0, h.Len())
	n, err := h.Pop()
	assert.Equal(t, ErrEmptyQueue, err)
	assert.Equal(t, 0, n)
}
