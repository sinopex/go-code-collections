// 118. 杨辉三角
// https://leetcode-cn.com/problems/pascals-triangle/
package array

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1 // 每一行行首为0
		if i >= 1 {
			result[i][i] = 1         // 从第二行开始，每行行尾为0
			for j := 1; j < i; j++ { // 中间每个值，等于上一行的两个值相加
				result[i][j] = result[i-1][j] + result[i-1][j-1]
			}
		}
	}

	return result
}
