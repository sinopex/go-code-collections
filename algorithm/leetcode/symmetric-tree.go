// 101 对称二叉树
// https://leetcode-cn.com/problems/symmetric-tree/
package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return _isSymmetric(root.Left, root.Right)
}

func _isSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && _isSymmetric(left.Left, right.Right) && _isSymmetric(left.Right, right.Left)
}
