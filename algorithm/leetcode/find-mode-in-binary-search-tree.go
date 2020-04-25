// 501 二叉搜索树种的众数
// https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/
package leetcode

import (
	"math"
)

func findMode(root *TreeNode) []int {
	var values = make(map[int]int)
	_findMode(root, values)
	var num []int

	max := -math.MaxInt64

	for _, v := range values {
		if v > max {
			max = v
		}
	}

	for k, v := range values {
		if v == max {
			num = append(num, k)
		}
	}

	return num
}

func _findMode(root *TreeNode, values map[int]int) {
	if root == nil {
		return
	}

	values[root.Val] += 1
	_findMode(root.Left, values)
	_findMode(root.Right, values)
}
