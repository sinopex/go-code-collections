// 674. 最长连续递增序列
// https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/
package array

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev, curLength, maxLength := nums[0], 1, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > prev {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			curLength = 1
		}
		prev = nums[i]
	}
	if curLength > maxLength {
		maxLength = curLength
	}

	return maxLength
}
