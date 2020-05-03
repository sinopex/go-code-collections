// 面试题 01.09. 字符串轮转1
// https://leetcode-cn.com/problems/string-rotation-lcci/
package leetcode

import (
	"strings"
)

func isFlipedString(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 != n2 {
		return false
	}

	for i := n2 - 1; i >= 0; i-- {
		if s2[i:n2]+s2[0:i] == s1 {
			return true
		}
	}

	return false
}

// 借鉴leetcode大牛的写法，直接相加，求匹配
func isFlipedString2(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s2+s2, s1)
}
