// 766. 托普利茨矩阵
// https://leetcode-cn.com/problems/toeplitz-matrix/
package array

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

func isToeplitzMatrix3(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	table := make(map[int]int, m+n)
	// 第一列
	for i := 0; i < m; i++ {
		table[100+i] = matrix[i][0]
	}

	// 第一行
	for i := 1; i < n; i++ {
		table[i] = matrix[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			key := 0
			if i >= j {
				key = (i - j) + 100
			} else {
				key = j - i
			}
			if table[key] != matrix[i][j] {
				return false
			}
		}
	}

	return true
}
