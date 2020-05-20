// 1051. 高度检查器
// https://leetcode-cn.com/problems/height-checker/
package array

import (
	"sort"
)

func heightChecker(heights []int) int {
	n := len(heights)
	if n <= 1 {
		return n
	}
	var sortedHeights = make([]int, n)
	copy(sortedHeights, heights)
	sort.Ints(sortedHeights)

	var count int
	for i := 0; i < n; i++ {
		if sortedHeights[i] != heights[i] {
			count++
		}
	}

	return count
}
