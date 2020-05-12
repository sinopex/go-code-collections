// 521. 最长特殊序列 Ⅰ
// https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/
package string

func findLUSlength(a string, b string) int {
	na, nb := len(a), len(b)
	if na > nb {
		return na
	}

	if na < nb {
		return nb
	}

	if a != b {
		return na
	}

	return -1
}
