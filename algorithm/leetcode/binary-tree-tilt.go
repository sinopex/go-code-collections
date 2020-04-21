// 563 二叉树的坡度
// https://leetcode-cn.com/problems/binary-tree-tilt/
package leetcode

func findTilt(root *TreeNode) int {
	var n int
	_findTilt(root, &n)
	return n
}

func _findTilt(root *TreeNode, n *int) int {
	if root == nil {
		return 0
	}
	left := _findTilt(root.Left, n)
	right := _findTilt(root.Right, n)

	if left > right {
		*n += left - right
	} else {
		*n += right - left
	}

	return left + right + root.Val
}
