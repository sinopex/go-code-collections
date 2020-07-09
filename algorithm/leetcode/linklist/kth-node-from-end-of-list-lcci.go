// 面试题 02.02. 返回倒数第 k 个节点
// https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
package linklist

func kthToLast(head *ListNode, k int) int {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow.Val
}
