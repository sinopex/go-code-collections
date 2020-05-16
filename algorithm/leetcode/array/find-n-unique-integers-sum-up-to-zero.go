// 1304. 和为零的N个唯一整数
// https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero/
package array

func sumZero(n int) []int {
	result := make([]int, n)

	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			result[i] = 0
			break
		}

		result[i] = i + 1
		result[j] = -1 * (i + 1)
	}

	return result
}
