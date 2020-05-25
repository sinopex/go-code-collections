// 1394. 找出数组中的幸运数
// https://leetcode-cn.com/problems/find-lucky-integer-in-an-array/
package array

func findLucky(arr []int) int {
	table := make(map[int]int, len(arr))
	lucky := -1

	for _, x := range arr {
		table[x]++
	}

	for k, v := range table {
		if k == v {
			if v > lucky {
				lucky = v
			}
		}
	}

	return lucky
}
