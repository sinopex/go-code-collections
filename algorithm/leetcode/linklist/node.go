package linklist

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

func (l *ListNode) String() string {
	var sb strings.Builder
	v := l
	dot := ""
	for v != nil {
		sb.WriteString(fmt.Sprintf("%s%d", dot, v.Val))
		dot = ","
		v = v.Next
	}
	return sb.String()
}
