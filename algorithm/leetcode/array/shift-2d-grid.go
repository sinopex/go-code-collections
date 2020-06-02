// 1260. 二维网格迁移
// https://leetcode-cn.com/problems/shift-2d-grid
package array

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	k %= m * n
	result := make([][]int, m)
	for i, _ := range result {
		result[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pos := i*n + j + k
			result[pos/n%m][pos%n] = grid[i][j]
		}
	}

	return result
}

func shiftGrid3(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	k %= m * n
	for i := 0; i < k; i++ {
		previous := grid[m-1][n-1]
		for j, _ := range grid {
			for k, _ := range grid[j] {
				previous, grid[j][k] = grid[j][k], previous
			}
		}
	}
	return grid
}

func shiftGrid2(grid [][]int, k int) [][]int {
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
