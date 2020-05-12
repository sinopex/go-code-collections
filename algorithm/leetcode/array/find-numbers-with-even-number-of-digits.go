// 1295. 统计位数为偶数的数字
// https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits/
package array

func findNumbers(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 10 {
			continue
		} else if nums[i] < 100 {
			result++
		} else if nums[i] < 1000 {
			continue
		} else if nums[i] < 10000 {
			result++
		} else if nums[i] < 100000 {
			continue
		} else {
		}
	}

	return result
}
