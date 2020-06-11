// 26. 删除排序数组中的重复项
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
package array

func removeDuplicates(nums []int) int {

	if n := len(nums); n <= 1 {
		return n
	}
	index, cursor, prev := 1, 1, nums[0]

	for cursor < len(nums) {
		if nums[cursor] != prev {
			nums[index] = nums[cursor]
			prev = nums[cursor]
			index++
		}
		cursor++
	}

	return index
}
