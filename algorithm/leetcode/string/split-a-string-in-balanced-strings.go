// 1221 分割平衡字符串
// https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
package string

func balancedStringSplit(s string) int {
	ln, rn, cnt := 0, 0, 0
	for _, c := range s {
		if c == 'L' {
			ln++
		}
		if c == 'R' {
			rn++
		}

		if ln == rn {
			cnt++
			ln, rn = 0, 0
		}
	}

	return cnt
}
