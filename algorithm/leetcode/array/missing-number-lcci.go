// 面试题 17.04. 消失的数字
// https://leetcode-cn.com/problems/missing-number-lcci/
package array

func missingNumber2(nums []int) int {
	n, c := len(nums), 0
	for _, v := range nums {
		c += v
	}

	return n*(1+n)/2 - c
}
func missingNumber(nums []int) int {
	var result int

	for k, v := range nums {
		result ^= k
		result ^= v
	}
	result ^= len(nums)

	return result
}
