// 1252. 奇数值单元格的数目
// https://leetcode-cn.com/problems/cells-with-odd-values-in-a-matrix/
package array

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	count := 0

	for _, set := range indices {
		// 行自增
		for i, _ := range matrix[set[0]] {
			matrix[set[0]][i] ^= 1 // 0,1 交替
		}

		// 列自增
		for _, v := range matrix {
			v[set[1]] ^= 1 // 0,1 交替
		}
	}

	for _, v := range matrix {
		for _, vv := range v {
			if vv == 1 {
				count++
			}
		}
	}

	return count
}
