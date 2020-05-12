// 709 转换成小写字母
// https://leetcode-cn.com/problems/to-lower-case/
package string

import (
	"strings"
)

func toLowerCase(str string) string {
	var sb strings.Builder

	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			c = c + 32
		}
		sb.WriteRune(c)
	}

	return sb.String()
}
