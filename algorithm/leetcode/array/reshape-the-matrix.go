// 566. 重塑矩阵
// https://leetcode-cn.com/problems/reshape-the-matrix/
package array

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	result := make([][]int, r)

	row, index := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if result[row] == nil {
				result[row] = make([]int, c)
			}
			result[row][index] = nums[i][j]
			index++
			if index == c {
				row++
				index = 0
			}
		}
	}

	return result
}

func matrixReshape2(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	result := make([][]int, r)
	index := 0
	queue, s := make([]int, c), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			queue[s] = nums[i][j]
			s++
			if s == c {
				result[index] = queue
				index++
				queue, s = make([]int, c), 0
			}
		}
	}

	return result
}
