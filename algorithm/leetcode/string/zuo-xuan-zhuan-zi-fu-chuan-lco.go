// 面试题58-II 左旋转字符串
// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
package string

import (
	"strings"
)

func reverseLeftWords(s string, n int) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	n = n % l

	var sb strings.Builder
	sb.WriteString(s[n:])
	sb.WriteString(s[0:n])
	return sb.String()
}
