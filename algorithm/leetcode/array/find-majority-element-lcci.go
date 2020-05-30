// 面试题 17.10. 主要元素
// https://leetcode-cn.com/problems/find-majority-element-lcci/
package array

func majorityElement(nums []int) int {
	major := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
			count = 1
		} else if nums[i] == major {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, v := range nums {
		if v == major {
			count++
		}
	}

	if count >= len(nums)/2+1 {
		return major
	}

	return -1
}
