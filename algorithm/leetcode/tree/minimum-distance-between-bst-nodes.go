// 783 二叉搜索树节点最小距离
// https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/
package tree

import (
	"math"
)

func minDiffInBST(root *TreeNode) int {
	var n, p int
	var flag bool
	n = math.MaxInt64
	_minDiffInBST(root, &n, &p, &flag)
	return n
}

func _minDiffInBST(root *TreeNode, n, p *int, flag *bool) {
	if root == nil {
		return
	}
	_minDiffInBST(root.Left, n, p, flag)

	if !*flag {
		// 第一个元素起点，不做比较
		*flag = true
	} else {
		//
		if root.Val-*p < *n {
			*n = root.Val - *p
		}
	}
	// 当前元素做后续元素，计算差值的对象
	*p = root.Val
	_minDiffInBST(root.Right, n, p, flag)
}
