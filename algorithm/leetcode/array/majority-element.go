// 169. 多数元素
// https://leetcode-cn.com/problems/majority-element
package array

func majorityElement1(nums []int) int {
	var major, count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
			count = 1
		} else if nums[i] == major {
			count++
		} else {
			count--
		}
	}

	return major
}
