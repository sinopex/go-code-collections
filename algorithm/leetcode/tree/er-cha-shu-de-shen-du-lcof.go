// 二叉树的深度
// https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
package tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
