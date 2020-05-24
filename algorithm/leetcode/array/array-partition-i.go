// 561. 数组拆分 I
// https://leetcode-cn.com/problems/array-partition-i/
package array

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var result int

	for i := 0; i < len(nums)-1; i += 2 {
		result += nums[i]
	}

	return result
}
