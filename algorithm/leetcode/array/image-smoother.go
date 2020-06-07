// 661. 图片平滑器
// https://leetcode-cn.com/problems/image-smoother/
package array

func imageSmoother(M [][]int) [][]int {
	m, n := len(M), len(M[0])
	result := make([][]int, m)
	for k, _ := range result {
		result[k] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, count := M[i][j], 1
			// 东
			if j+1 < n {
				sum += M[i][j+1]
				count++
			}
			// 南
			if i+1 < m {
				sum += M[i+1][j]
				count++
			}
			// 西
			if j-1 >= 0 {
				sum += M[i][j-1]
				count++
			}
			// 北
			if i-1 >= 0 {
				sum += M[i-1][j]
				count++
			}
			// 东南
			if i+1 < m && j+1 < n {
				sum += M[i+1][j+1]
				count++
			}
			// 东北
			if i-1 >= 0 && j+1 < n {
				sum += M[i-1][j+1]
				count++
			}
			// 西北
			if i-1 >= 0 && j-1 >= 0 {
				sum += M[i-1][j-1]
				count++
			}
			// 西南
			if i+1 < m && j-1 >= 0 {
				sum += M[i+1][j-1]
				count++
			}

			result[i][j] = sum / count
		}
	}

	return result
}
