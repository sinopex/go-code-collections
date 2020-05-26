// https://leetcode-cn.com/problems/fibonacci-number/
// https://leetcode-cn.com/problems/fibonacci-number/
package array

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}

	p1, p2 := 0, 1
	for i := 2; i <= N; i++ {
		p1, p2 = p2, p1+p2
	}
	return p2
}
