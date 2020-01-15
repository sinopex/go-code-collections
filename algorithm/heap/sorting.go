package heap

import (
	"sync"
)

// 通过构建最小堆实现堆排序

type Sorting struct {
	sync.Mutex
	len   int
	queue []int
}

func (s *Sorting) Push(n int) {
	s.Lock()
	defer s.Unlock()
	s.queue = append(s.queue, n)
	s.len += 1
	s.shiftUp(s.len - 1)
}
func (s *Sorting) Pop() (int, error) {
	s.Lock()
	defer s.Unlock()
	if s.len == 0 {
		return 0, ErrEmptyQueue
	}

	n := s.queue[0]               // 取出堆顶元素
	s.queue[0] = s.queue[s.len-1] // 将堆尾元素上任
	s.queue = s.queue[:s.len-1]   // 删除队尾元素
	s.len -= 1                    // 长度减1
	if s.len > 0 {
		s.shiftDown(0) // 堆顶元素重新分配
	}
	return n, nil
}
func (s *Sorting) Len() int {
	return s.len
}
func (s *Sorting) Peek() (int, error) {
	if s.len == 0 {
		return 0, ErrEmptyQueue
	}
	return s.queue[0], nil
}
func (s *Sorting) shiftDown(i int) {
	leftIndex, rightIndex := 2*i+1, 2*i+2
	// 左节点不为空
	if leftIndex <= s.len-1 {
		if s.queue[leftIndex] < s.queue[i] {
			s.queue[leftIndex], s.queue[i] = s.queue[i], s.queue[leftIndex]
			s.shiftDown(leftIndex)
		}
	}

	// 右节点不为空
	if rightIndex <= s.len-1 {
		if s.queue[rightIndex] < s.queue[i] {
			s.queue[rightIndex], s.queue[i] = s.queue[i], s.queue[rightIndex]
			s.shiftDown(rightIndex)
		}
	}
}
func (s *Sorting) shiftUp(i int) {
	if i == 0 {
		return
	}
	parentIndex := (i - 1) / 2
	if s.queue[parentIndex] > s.queue[i] {
		s.queue[parentIndex], s.queue[i] = s.queue[i], s.queue[parentIndex]
		s.shiftUp(parentIndex)
	}
}
