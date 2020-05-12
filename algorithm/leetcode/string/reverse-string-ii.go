// 541. 反转字符串 II
// https://leetcode-cn.com/problems/reverse-string-ii/
package string

import (
	"bytes"
)

func reverseStr(s string, k int) string {
	var buf bytes.Buffer
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		// 前半截 反序输出
		for j := i + k - 1; j >= i; j-- {
			if j < n {
				buf.WriteByte(s[j])

			}
		}

		// 后半截 正序输出
		for j := i + k; j < i+2*k; j++ {
			if j < n {
				buf.WriteByte(s[j])
			}
		}
	}

	return buf.String()
}
