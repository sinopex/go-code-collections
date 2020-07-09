// 面试题 02.01. 移除重复节点
// https://leetcode-cn.com/problems/remove-duplicate-node-lcci/
package linklist

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	table := map[int]bool{head.Val: true}
	pos := head
	for pos.Next != nil {
		n := pos.Next
		if !table[n.Val] {
			table[n.Val] = true
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}

	return head
}
