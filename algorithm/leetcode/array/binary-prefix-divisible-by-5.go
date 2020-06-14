// 1018. 可被 5 整除的二进制前缀
// https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/
package array

func prefixesDivBy5(A []int) []bool {
	count := 0
	result := make([]bool, len(A))
	for i := 0; i < len(A); i++ {
		count = (count<<1 + A[i]) % 5
		if count%5 == 0 {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}
