// 697. 数组的度
// https://leetcode-cn.com/problems/degree-of-an-array/
package array

func findShortestSubArray(nums []int) int {
	table := make(map[int][3]int)
	min, max := 50000, 0
	for k, v := range nums {
		if _, ok := table[v]; !ok {
			table[v] = [3]int{k, k, 1}
		} else {
			table[v] = [3]int{table[v][0], k, table[v][2] + 1}
		}
		if table[v][2] > max {
			max = table[v][2]
		}
	}

	for _, v := range table {
		if v[2] == max {
			if v[1]-v[0]+1 < min {
				min = v[1] - v[0] + 1
			}
		}
	}

	return min
}
