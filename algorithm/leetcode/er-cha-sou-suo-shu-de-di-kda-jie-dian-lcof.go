// 面试题54 二叉搜索树的第K大节点
// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
package leetcode

func kthLargest(root *TreeNode, k int) int {
	if root == nil || k == 0 {
		return 0
	}

	h := NewMinHeap(k)
	_kthLargest(root, h)
	return h.Top()
}

func _kthLargest(root *TreeNode, h *MinHeap) {
	if root == nil {
		return
	}

	_kthLargest(root.Left, h)
	h.Push(root.Val)
	_kthLargest(root.Right, h)
}
