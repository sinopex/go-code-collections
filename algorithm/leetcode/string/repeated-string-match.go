// 686. 重复叠加字符串匹配
// https://leetcode-cn.com/problems/repeated-string-match/
package string

import (
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	L := len(B) / len(A)

	for i := 0; i <= 2; i++ {
		if strings.Contains(strings.Repeat(A, L+i), B) {
			return L + i
		}
	}

	return -1
}
