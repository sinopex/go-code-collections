// 110 平衡二叉树
// https://leetcode-cn.com/problems/balanced-binary-tree/submissions/
package tree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	_, b := _isBalancedDepth(root)
	return b
}

func _isBalancedDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	ld, lt := _isBalancedDepth(root.Left)
	rd, rt := _isBalancedDepth(root.Right)

	depth := ld
	x := ld - rd
	if rd > ld {
		depth = rd
		x = rd - ld
	}
	depth++

	return depth, lt && rt && x <= 1
}
