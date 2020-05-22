// 977. 有序数组的平方
// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
package array

import (
	"math"
	"sort"
)

func sortedSquares(A []int) []int {
	n := len(A)
	result := make([]int, n)
	// 根据绝对值排序
	sort.Slice(A, func(i, j int) bool {
		return math.Abs(float64(A[i])) < math.Abs(float64(A[j]))
	})

	// 从小到大写入
	for i, v := range A {
		result[i] = v * v
	}

	return result
}
