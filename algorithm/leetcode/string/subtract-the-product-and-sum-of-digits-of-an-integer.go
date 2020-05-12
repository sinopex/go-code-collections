// 1281 整数的各位积和之差
// https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
package string

func subtractProductAndSum(n int) int {
	if n == 0 {
		return 0
	}
	sum, mul := 0, 1

	for n > 0 {
		x := n % 10
		sum += x
		mul *= x
		n = n / 10
	}

	return mul - sum
}
