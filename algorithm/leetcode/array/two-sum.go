// 1. 两数之和
// https://leetcode-cn.com/problems/two-sum/submissions/
package array

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k2, v := range nums {
		if k1, ok := m[target-v]; ok {
			return []int{k1, k2}
		} else {
			m[v] = k2
		}
	}

	return nil
}
