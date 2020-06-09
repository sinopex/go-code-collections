// 896. 单调数列
// https://leetcode-cn.com/problems/monotonic-array/
package array

func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	var flag int // 0未初始化 1递增 2 递减

	for i := 0; i < len(A)-1; i++ {

		if A[i] == A[i+1] {
			continue
		}

		if A[i] > A[i+1] {
			if flag == 0 {
				flag = 2
			} else if flag == 1 {
				return false
			} else {
				continue
			}
		}

		if A[i] < A[i+1] {
			if flag == 0 {
				flag = 1
			} else if flag == 2 {
				return false
			} else {
				continue
			}
		}
	}
	return true
}
