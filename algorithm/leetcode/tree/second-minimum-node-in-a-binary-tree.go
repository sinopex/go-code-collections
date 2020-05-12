// 671 二叉树中第二小的节点
// https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/
package tree

func findSecondMinimumValue(root *TreeNode) int {
	var n int = -1
	_findSecondMinimumValue(root, root.Val, &n)
	return n
}
func _findSecondMinimumValue(root *TreeNode, b int, n *int) {
	if root == nil {
		return
	}

	// 还没赋值
	if *n == -1 && root.Val != b {
		*n = root.Val
	}

	// 找到一个与根节点不一样的，且比当前第二小还小的数值
	if *n > root.Val && root.Val != b {
		*n = root.Val
	}

	_findSecondMinimumValue(root.Left, b, n)
	_findSecondMinimumValue(root.Right, b, n)
}
