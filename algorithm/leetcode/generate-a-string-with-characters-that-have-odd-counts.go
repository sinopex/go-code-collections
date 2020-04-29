// 1374 生成每种字符都是奇数个的字符串
// https://leetcode-cn.com/problems/generate-a-string-with-characters-that-have-odd-counts/
package leetcode

import (
	"strings"
)

func generateTheString(n int) string {
	if n&1 == 1 {
		return strings.Repeat("a", n)
	}

	return strings.Repeat("a", n-1) + "b"
}
