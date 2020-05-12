// 面试题 01.04. 回文排列
// https://leetcode-cn.com/problems/palindrome-permutation-lcci/
package string

func canPermutePalindrome(s string) bool {
	var buf = []byte(s)
	var dict = make(map[byte]int)
	var oddNum = 0
	for _, v := range buf {
		if dict[v] > 0 {
			dict[v]--
		} else {
			dict[v]++
		}
	}

	for _, v := range dict {
		if v == 1 {
			oddNum++
		}
		if oddNum >= 2 {
			return false
		}
	}

	return true
}
