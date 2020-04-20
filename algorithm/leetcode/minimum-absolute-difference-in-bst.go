// 530 二叉搜索树的最小绝对差
// minimum-absolute-difference-in-bst
package leetcode

import (
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	var n, p int
	var flag bool
	n = math.MaxInt64
	_getMinimumDifference(root, &n, &p, &flag)
	return n
}

func _getMinimumDifference(root *TreeNode, n, p *int, flag *bool) {
	if root == nil {
		return
	}
	_getMinimumDifference(root.Left, n, p, flag)

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
	_getMinimumDifference(root.Right, n, p, flag)
}
