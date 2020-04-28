// LCP01 猜数字
// https://leetcode-cn.com/problems/guess-numbers/
package leetcode

func game(guess []int, answer []int) int {
	if len(guess) == 0 || len(answer) == 0 {
		return 0
	}
	total := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			total++
		}
	}
	return total
}
