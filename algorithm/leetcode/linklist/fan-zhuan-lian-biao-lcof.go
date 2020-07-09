// 剑指 Offer 24. 反转链表
// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
package linklist

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		head.Next, result, head = result, head, head.Next
	}

	return result
}
