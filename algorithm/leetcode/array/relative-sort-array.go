// 1122. 数组的相对排序
// https://leetcode-cn.com/problems/relative-sort-array/
package array

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	table := make(map[int]int, len(arr2))
	// 将相对顺序的数字转换成哈希表
	for k, v := range arr2 {
		table[v] = k
	}

	sort.Slice(arr1, func(i, j int) bool {
		index1, ok1 := table[arr1[i]]
		index2, ok2 := table[arr1[j]]

		if ok1 && ok2 {
			return index1 < index2
		} else if ok1 {
			return true
		} else if ok2 {
			return false
		} else {
			return arr1[i] < arr1[j]
		}
	})

	return arr1
}
