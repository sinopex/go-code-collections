// 832. 翻转图像
// https://leetcode-cn.com/problems/flipping-an-image/
package array

func flipAndInvertImage(A [][]int) [][]int {
	columns := len(A[0])

	for rowIndex, rowData := range A {
		for i, j := 0, columns-1; i <= j; i, j = i+1, j-1 {
			A[rowIndex][i], A[rowIndex][j] = rowData[j]^1, rowData[i]^1
		}
	}

	return A
}
