// 面试题 10.01. 合并排序的数组
// https://leetcode-cn.com/problems/sorted-merge-lcci/
package array

func merge(A []int, m int, B []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if A[i] > B[j] {
			A[k] = A[i]
			i--
		} else {
			A[k] = B[j]
			j--
		}
		k--
	}

	for i >= 0 {
		A[k] = A[i]
		k--
		i--
	}

	for j >= 0 {
		A[k] = B[j]
		k--
		j--
	}

}
