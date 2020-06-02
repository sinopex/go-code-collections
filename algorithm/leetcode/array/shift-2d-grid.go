// 1260. 二维网格迁移
// https://leetcode-cn.com/problems/shift-2d-grid
package array

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	all := make([]int, m*n)
	for index, row := range grid {
		copy(all[index*n:(index+1)*n], row)
	}
	k %= m * n
	all = append(all[m*n-k:], all[0:m*n-k]...)
	for index, row := range grid {
		copy(row, all[index*n:(index+1)*n])
	}

	return grid
}
