// 面试题03. 数组中重复的数字
// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
package array

func findRepeatNumber(nums []int) int {
	table := make(map[int]bool, len(nums))

	for _, v := range nums {
		if table[v] {
			return v
		}
		table[v] = true
	}

	return 0
}
