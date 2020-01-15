package heap

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const defaultMaxSize = 100

func mockMessage(priority int) Message {
	// 模拟是订单业务，每个订单有个优先级
	// 优先处理Priority值大的订单
	return Message{
		Priority: priority,
		Body:     []byte(fmt.Sprintf(`{"trade_no":%s}`, fmt.Sprintf("%s", time.Now().Format("20060102150406")))),
	}
}

func TestPriorityQueue_Create(t *testing.T) {
	pq := NewPriorityQueue(defaultMaxSize)
	assert.Equal(t, 0, pq.Len())
}

func TestPriorityQueue_Insert(t *testing.T) {
	pq := NewPriorityQueue(defaultMaxSize)
	// 将模拟数据建堆
	dataset := []int{15, 13, 18, 17, 14, 10, 13}
	for _, v := range dataset {
		assert.Nil(t, pq.Push(mockMessage(v)))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(dataset)))
	n := len(dataset)
	// 依次将堆的数据取出
	for _, v := range dataset {
		assert.Equal(t, n, pq.len)
		peek, err := pq.Peek()
		fmt.Printf("%d,", peek.Priority)
		assert.Nil(t, err)
		assert.Equal(t, v, peek.Priority)

		message, err := pq.Pop()
		assert.Nil(t, err)
		assert.Equal(t, v, message.Priority)
		n--
	}
	fmt.Println()
}
