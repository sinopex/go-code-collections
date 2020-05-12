// 1365. 有多少小于当前数字的数字
// https://leetcode-cn.com/problems/how-many-numbers-are-s
package array

func smallerNumbersThanCurrent(nums []int) []int {
	var sum [101]int

	for _, v := range nums {
		sum[v]++
	}

	for i := 1; i < len(sum); i++ {
		sum[i] += sum[i-1]
	}

	result := make([]int, len(nums))
	for i, n := range nums {
		if n > 0 {
			result[i] = sum[n-1]
		}
	}
	return result
}
