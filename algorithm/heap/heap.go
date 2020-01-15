package heap

import (
	"errors"
)

var (
	nilMessage        = Message{}
	ErrLengthExceeded = errors.New("heap.PriorityQueue: length exceeded")
	ErrEmptyQueue     = errors.New("heap.PriorityQueue: empty queue")
)
