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

func sortedSquares2(A []int) []int {
	n := len(A)
	result := make([]int, n)
	for i, j, k := 0, n-1, n-1; i <= j; {
		// 左边大
		if math.Abs(float64(A[i])) >= math.Abs(float64(A[j])) {
			result[k] = A[i] * A[i]
			i++
		} else {
			result[k] = A[j] * A[j]
			j--
		}
		k--
	}

	return result
}
