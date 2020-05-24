// 1413. 逐步求和得到正数的最小值
// https://leetcode-cn.com/problems/minimum-value-to-get-positive-step-by-step-sum/
package array

func minStartValue(nums []int) int {
	sum, min := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if sum < min {
			min = sum
		}
	}

	if min >= 0 {
		return 1
	}

	return -min + 1
}
