// 1380. 矩阵中的幸运数
// https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
package array

func luckyNumbers(matrix [][]int) []int {
	result, h := make([]int, 0), make(map[int]bool)
	m, n := len(matrix), len(matrix[0])

	var getMin = func(arr []int) int {
		min := arr[0]
		for i := 1; i < n; i++ {
			if arr[i] < min {
				min = arr[i]
			}

		}
		return min
	}

	// 每一行的最小值写入Hash
	for i := 0; i < m; i++ {
		h[getMin(matrix[i])] = true
	}

	for i := 0; i < n; i++ {
		// 找出每一列的最大值，如果最大值存在最小值的Hash中
		max := matrix[0][i]
		for j := 1; j < m; j++ {
			if matrix[j][i] > max {
				max = matrix[j][i]
			}
		}

		// 则满足目标
		if h[max] {
			result = append(result, max)
		}
	}

	return result
}
