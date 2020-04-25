// 501 二叉搜索树种的众数
// https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/
package leetcode

import (
	"math"
)

func findMode(root *TreeNode) []int {
	values := _findMode(root)
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

func _findMode(root *TreeNode) map[int]int {
	if root == nil {
		return map[int]int{}
	}

	values := make(map[int]int)
	values[root.Val] = 1
	for n, v := range _findMode(root.Left) {
		values[n] += v
	}

	for n, v := range _findMode(root.Right) {
		values[n] += v
	}

	return values
}
