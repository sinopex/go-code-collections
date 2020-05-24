// 905. 按奇偶排序数组
// https://leetcode-cn.com/problems/sort-array-by-parity/
package array

func sortArrayByParity2(A []int) []int {
	n := len(A)
	for pos, search := 0, 0; search < n; {
		// 如果当前是偶数，则继续往前走
		if A[pos]&1 == 0 {
			pos++
			if search < pos {
				search = pos
			}
			continue
		}

		// 如果是奇数，则一直往后查询
		for search < n && A[search]&1 == 1 {
			search++
		}

		// 成功找到一个偶数
		// 将这个偶数放到当前偶数序列的最后
		if search < n && A[search]&1 == 0 {
			A[pos], A[search] = A[search], A[pos]
		}
	}
	return A
}
func sortArrayByParity(A []int) []int {
	index := 0
	for i := 0; i < len(A); i++ {
		if A[i]&1 == 0 {
			A[i], A[index] = A[index], A[i]
			index++
		}
	}

	return A
}

// func sortArrayByParity(A []int) []int {
// 	i, j := 0, len(A)-1
//
// 	for i < j {
// 		if A[i]&1 == 1 && A[j]&1 == 0 {
// 			A[i], A[j] = A[j], A[i]
// 			i++
// 			j--
// 		}else if A[i]&1==0{
// 			i++
// 		}else if A[j]&1==1{
// 			j--
// 		}
// 	}
// 	return A
// }
