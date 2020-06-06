// 485. 最大连续1的个数
// https://leetcode-cn.com/problems/max-consecutive-ones/
package array

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			i++
			continue
		}

		j := i
		for ; j < len(nums); j++ {
			if nums[j] == 0 {
				break
			}
		}

		if j-i > max {
			max = j - i
		}
		i = j
	}

	return max
}
