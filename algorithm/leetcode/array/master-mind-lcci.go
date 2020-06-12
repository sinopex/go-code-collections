// 面试题 16.15. 珠玑妙算
// https://leetcode-cn.com/problems/master-mind-lcci/
package array

func masterMind(solution string, guess string) []int {
	result := make([]int, 2)
	s, g := make(map[byte]int), make(map[byte]int)
	for i := 0; i < 4; i++ {
		if guess[i] == solution[i] {
			result[0]++
		} else {
			s[solution[i]]++
			g[guess[i]]++
		}
	}

	var min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for x, v := range s {
		result[1] += min(v, g[x])
	}

	return result
}
