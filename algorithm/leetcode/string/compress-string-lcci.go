// 面试题 01.06. 字符串压缩
// https://leetcode-cn.com/problems/compress-string-lcci/
package string

import (
	"bytes"
	"strconv"
)

func compressString(S string) string {
	var s = []byte(S)
	var ret bytes.Buffer
	var cnt int
	n := len(s)

	for i := 0; i < n; {
		cnt = 1
		for j := i + 1; j < n; j++ {
			if s[j] == s[i] {
				cnt++
			} else {
				break
			}
		}
		ret.WriteByte(s[i])
		ret.WriteString(strconv.Itoa(cnt))
		i += cnt
	}

	if ret.Len() >= n {
		return S
	}

	return ret.String()
}
