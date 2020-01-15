package heap

import (
	"sync"
)

// 通过构建最大堆实现优先队列

type Message struct {
	Priority int
	Body     []byte
}

func NewPriorityQueue(cap int) *PriorityQueue {
	return &PriorityQueue{
		cap:   cap,
		len:   0,
		queue: make([]Message, 0, cap),
	}
}

type PriorityQueue struct {
	sync.Mutex
	cap   int // 堆的最大容量
	len   int // 对的当前容量
	queue []Message
}

func (p *PriorityQueue) Len() int {
	return p.len
}

func (p *PriorityQueue) Peek() (Message, error) {
	if p.len == 0 {
		return nilMessage, ErrEmptyQueue
	}
	return p.queue[0], nil
}

func (p *PriorityQueue) Push(message Message) error {
	p.Lock()
	defer p.Unlock()
	if p.len == p.cap {
		return ErrLengthExceeded
	}
	p.queue = append(p.queue, message)
	p.len += 1
	p.shiftUp(p.len - 1)
	return nil
}

func (p *PriorityQueue) Pop() (Message, error) {
	if p.len == 0 {
		return nilMessage, ErrEmptyQueue
	}
	p.Lock()
	defer p.Unlock()
	message := p.queue[0]

	p.queue[0] = p.queue[p.len-1] // 将堆尾元素压入堆顶
	p.queue = p.queue[:p.len-1]   // 删除队尾元素
	p.len -= 1
	if p.len > 1 {
		p.shiftDown(0)
	}

	return message, nil
}

func (p *PriorityQueue) shiftUp(i int) {
	// 已经到堆顶了
	if i == 0 {
		return
	}

	parentIndex := (i - 1) / 2
	if p.queue[parentIndex].Priority < p.queue[i].Priority {
		p.queue[parentIndex], p.queue[i] = p.queue[i], p.queue[parentIndex]
		p.shiftUp(parentIndex)
	}
}

func (p *PriorityQueue) shiftDown(i int) {
	leftIndex, rightIndex := i*2+1, i*2+2

	if leftIndex <= p.len-1 {
		if p.queue[leftIndex].Priority > p.queue[i].Priority {
			p.queue[leftIndex], p.queue[i] = p.queue[i], p.queue[leftIndex]
			p.shiftDown(leftIndex)
		}
	}

	if rightIndex <= p.len-1 {
		if p.queue[rightIndex].Priority > p.queue[i].Priority {
			p.queue[rightIndex], p.queue[i] = p.queue[i], p.queue[rightIndex]
			p.shiftDown(rightIndex)
		}
	}
}
