// 35. 搜索插入位置
// https://leetcode-cn.com/problems/search-insert-position/
package array

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}
