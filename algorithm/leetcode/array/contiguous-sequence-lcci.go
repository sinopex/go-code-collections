// 面试题 16.17. 连续数列
// https://leetcode-cn.com/problems/contiguous-sequence-lcci/
package array

import (
	"math"
)

func maxSubArray(nums []int) int {
	sum, max := 0, -math.MaxInt32
	for _, v := range nums {
		if sum+v <= v {
			sum = v
		} else {
			sum += v
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
func maxSubArray1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var sum int
	for i := 0; i < n; i++ {
		if i == 0 {
			sum = nums[i]
		}
		var tempSum int
		for j := i; j < n; j++ {
			if j == i {
				tempSum = nums[j]
			} else if tempSum+nums[j] < 0 { // 遇到和是小数则退出
				break
			} else {
				tempSum += nums[j]
			}
			if tempSum > sum {
				sum = tempSum
			}
		}
	}

	return sum
}
