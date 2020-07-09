package linklist

import (
	"encoding/json"
	"fmt"
	"testing"
)

func NewNodeList(a ...int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	v := &ListNode{
		Val:  a[0],
		Next: NewNodeList(a[1:]...),
	}
	return v
}
func Test(t *testing.T) {
	b, _ := json.MarshalIndent(NewNodeList(1, 2, 3, 4, 5), "", "\t")
	fmt.Printf("%s\n", string(b))
}
