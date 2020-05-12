package tree

import (
	"sync"
)

type MinHeap struct {
	Data []int
	Size int
	sync.Mutex
	length int
}

func (m *MinHeap) Top() int {
	m.Lock()
	defer m.Unlock()
	if len(m.Data) > 0 {
		return m.Data[0]
	}
	return 0
}

func (m *MinHeap) Push(val int) {
	m.Lock()
	defer m.Unlock()
	// 堆已满
	if len(m.Data) == m.Size {
		if val > m.Data[0] {
			m.Data[0] = val
			m.shiftDown(0)
		}
		return
	}

	m.Data = append(m.Data, val)
	m.length += 1
	m.shiftUp(m.length - 1)

	return
}

func (m *MinHeap) Pop() int {
	m.Lock()
	defer m.Unlock()
	if m.length == 0 {
		return 0
	}
	val := m.Data[0]
	m.Data[0] = m.Data[m.length-1]
	m.Data = m.Data[:m.length-1]
	m.length--

	return val
}

func (m *MinHeap) shiftUp(i int) {
	// 已经到堆顶了
	if i == 0 {
		return
	}

	parentIndex := (i - 1) / 2
	if m.Data[parentIndex] > m.Data[i] {
		m.Data[parentIndex], m.Data[i] = m.Data[i], m.Data[parentIndex]
		m.shiftUp(parentIndex)
	}
}

func (m *MinHeap) shiftDown(i int) {
	leftIndex, rightIndex := i*2+1, i*2+2

	if leftIndex <= m.length-1 {
		if m.Data[leftIndex] < m.Data[i] {
			m.Data[leftIndex], m.Data[i] = m.Data[i], m.Data[leftIndex]
			m.shiftDown(leftIndex)
		}
	}

	if rightIndex <= m.length-1 {
		if m.Data[rightIndex] < m.Data[i] {
			m.Data[rightIndex], m.Data[i] = m.Data[i], m.Data[rightIndex]
			m.shiftDown(rightIndex)
		}
	}
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		Size: size,
		Data: make([]int, 0),
	}
}
