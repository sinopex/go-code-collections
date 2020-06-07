// 167. 两数之和 II - 输入有序数组
// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
package array

func twoSum(numbers []int, target int) []int {
	for left, right := 0, len(numbers)-1; left < right; {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum > target {
			right--
		}

		if sum < target {
			left++
		}
	}

	return nil
}
