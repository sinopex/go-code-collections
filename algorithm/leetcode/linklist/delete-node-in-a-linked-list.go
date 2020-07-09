// 237. 删除链表中的节点
// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
package linklist

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
