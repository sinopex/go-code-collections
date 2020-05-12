// 404 左叶子之和
// https://leetcode-cn.com/problems/sum-of-left-leaves/
package tree

func sumOfLeftLeaves(root *TreeNode) int {
	var n int
	_sumOfLeftLeaves(root, &n)
	return n
}

func _sumOfLeftLeaves(root *TreeNode, n *int) {
	if root == nil {
		return
	}

	// 当前元素的左节点不为空，且左节点为叶子节点
	// 则元素的左节点满足要求
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*n += root.Left.Val
	}

	_sumOfLeftLeaves(root.Left, n)
	_sumOfLeftLeaves(root.Right, n)
}
