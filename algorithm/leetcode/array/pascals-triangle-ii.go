// 119. 杨辉三角 II
// https://leetcode-cn.com/problems/pascals-triangle-ii/
package array

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	index := 0
	for i := 0; i <= rowIndex; i++ {
		for k := index - 1; k > 0; k-- {
			row[k] = row[k] + row[k-1]
		}

		// 末尾为1
		row[index] = 1
		index++
	}

	return row
}
