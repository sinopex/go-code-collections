// 1342. 将数字变成0的操作次数
// https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/
package string

func numberOfSteps(num int) int {
	var total int
	for num > 0 {
		if num&1 == 1 {
			num--
		} else {
			num /= 2
		}
		total++
	}
	return total
}
