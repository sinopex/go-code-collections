// 1351. 统计有序矩阵中的负数
// https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/
package array

func countNegatives(grid [][]int) int {
	num, m, pos := 0, len(grid[0]), len(grid[0])-1

	for _, x := range grid {
		var i int
		for i = pos; i >= 0; i-- {
			if x[i] >= 0 {
				if i+1 < m {
					pos = i + 1
					num += m - pos
				}
				break
			}
		}

		if i == -1 {
			num += m
			pos = -1
		}
	}

	return num
}
