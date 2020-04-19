// 965 单值二叉树
// https://leetcode-cn.com/problems/univalued-binary-tree/
package leetcode

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return _isUnivalTree(root, root.Val)
}

func _isUnivalTree(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}

	return _isUnivalTree(root.Left, val) && root.Val == val && _isUnivalTree(root.Right, val)
}
