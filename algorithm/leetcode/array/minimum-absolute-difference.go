// 1200. 最小绝对差
// https://leetcode-cn.com/problems/minimum-absolute-difference/
package array

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	var result = make([][]int, 0)
	var min = math.MaxInt32

	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		temp := arr[i] - arr[i-1]
		if temp == min {
			result = append(result, []int{arr[i-1], arr[i]})
		} else if temp < min {
			min = temp
			result = [][]int{{arr[i-1], arr[i]}}
		}
	}

	return result
}
