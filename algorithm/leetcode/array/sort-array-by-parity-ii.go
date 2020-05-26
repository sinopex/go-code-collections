// 922. 按奇偶排序数组 II
// https://leetcode-cn.com/problems/sort-array-by-parity-ii/
package array

func sortArrayByParityII(A []int) []int {
	n := len(A)
	for i := 0; i < n-1; i += 2 {
		// 如果偶数位是奇数，则从后续找出一个偶数，与其替换
		if A[i]&1 == 1 {
			for j := i + 1; j < n; j++ {
				if A[j]&1 == 0 {
					A[i], A[j] = A[j], A[i]
					break
				}
			}
		}

		// 如果奇数位是偶数，则从后续找出一个奇数，与其替换
		if A[i+1]&1 == 0 {
			for j := i + 2; j < n; j++ {
				if A[j]&1 == 1 {
					A[i+1], A[j] = A[j], A[i+1]
					break
				}
			}
		}
	}

	return A
}

// 膜拜一个大神的写法
func sortArrayByParityII2(A []int) []int {
	left, right := 0, 1 // 奇数 偶数两个一组，往后排
	for left < len(A)-1 {
		// 偶数位的值不是偶数
		for A[left]%2 != 0 {
			// 则从奇数位上，找到一个偶数值，与其替换
			if A[right]%2 == 0 {
				A[left], A[right] = A[right], A[left]
			} else {
				right += 2
			}
		}
		left += 2
	}
	return A
}
