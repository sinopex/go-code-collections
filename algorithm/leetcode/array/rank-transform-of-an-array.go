// 1331. 数组序号转换
// https://leetcode-cn.com/problems/rank-transform-of-an-array/
package array

import (
	"sort"
)

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	m := make(map[int]int)
	start := 1
	m[sorted[0]] = start
	for i := 1; i < len(arr); i++ {
		if sorted[i] == sorted[i-1] {
			continue
		}
		start++
		m[sorted[i]] = start
	}

	for k, v := range arr {
		arr[k] = m[v]
	}

	return arr
}
