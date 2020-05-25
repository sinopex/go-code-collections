// 867. 转置矩阵
// https://leetcode-cn.com/problems/transpose-matrix/
package array

func transpose(A [][]int) [][]int {
	m := len(A)
	n := len(A[0])
	r := make([][]int, n)

	for i := range A {
		for j := range A[i] {
			if r[j] == nil {
				r[j] = make([]int, m)
			}
			r[j][i] = A[i][j]
		}
	}

	return r
}
