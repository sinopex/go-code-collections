package linklist

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
