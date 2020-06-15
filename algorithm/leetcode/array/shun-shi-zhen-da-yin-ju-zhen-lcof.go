// 面试题29. 顺时针打印矩阵
// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
package array

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	columns, rows := len(matrix[0]), len(matrix)
	minColumn, maxColumn := 0, columns-1
	minRow, maxRow := 0, rows-1
	result := make([]int, columns*rows)
	count := 0

	for minColumn <= maxColumn && minRow <= maxRow {
		// fmt.Printf("minColumn=%d,maxColumn=%d,minRow=%d,maxRow=%d\n", minColumn, maxColumn, minRow, maxRow)

		// 写入第一行
		for x := minColumn; x <= maxColumn; x++ {
			result[count] = matrix[minRow][x]
			count++
		}

		// 写入第一行与最后一行之间行的右侧尾值
		if maxRow-minRow > 1 {
			for x := minRow + 1; x < maxRow; x++ {
				result[count] = matrix[x][maxColumn]
				count++
			}
		}
		// 写入最后一行
		if minRow != maxRow {
			for k := maxColumn; k >= minColumn; k-- {
				result[count] = matrix[maxRow][k]
				count++
			}
		}

		// 写入第一行与最后一行之间行的左侧开始值
		if maxRow-minRow > 1 && minColumn != maxColumn {
			for x := maxRow - 1; x > minRow; x-- {
				result[count] = matrix[x][minColumn]
				count++
			}
		}

		minRow++
		maxRow--
		minColumn++
		maxColumn--

	}

	return result
}
