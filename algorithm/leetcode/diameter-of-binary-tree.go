// 543 二叉树的直径
// https://leetcode-cn.com/problems/diameter-of-binary-tree/
package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	_ = _diameterOfBinaryTree(root, &max)
	return max
}

func _diameterOfBinaryTree(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	l := _diameterOfBinaryTree(root.Left, max)
	r := _diameterOfBinaryTree(root.Right, max)

	// 如果左节点 + 右节点 + 父节点，组成的树高度超过最大值
	// 则更新为新的最大值
	if l+r > *max {
		*max = l + r
	}

	// 左边直径高，加上父节点，就是当前节点的最长直径
	if l > r {
		return l + 1
	}

	// 否则，就是右节点
	return r + 1
}
