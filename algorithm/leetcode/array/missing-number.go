// 268. 缺失数字
// https://leetcode-cn.com/problems/missing-number/
package array

func missingNumber3(nums []int) int {
	var result int
	// total=n(n+1)/2
	for i := 0; i < len(nums); i++ {
		result = result ^ i ^ nums[i]
	}

	return result ^ len(nums)
}
