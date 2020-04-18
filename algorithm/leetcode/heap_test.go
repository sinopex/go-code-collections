package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHeap(t *testing.T) {
	h := NewMinHeap(10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		value := r.Intn(100)
		h.Push(value)
		fmt.Printf("push=%d,top=%d\n", value, h.Top())
	}
}
