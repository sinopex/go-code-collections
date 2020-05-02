// 824. 山羊拉丁文
// https://leetcode-cn.com/problems/goat-latin/
package leetcode

import (
	"strings"
)

func toGoatLatin(S string) string {
	var sb strings.Builder
	cnt := 1
	var beginningByte byte
	var beginning = true
	var d = map[byte]bool{
		'a': true,
		'A': true,
		'e': true,
		'E': true,
		'i': true,
		'I': true,
		'o': true,
		'O': true,
		'u': true,
		'U': true,
	}
	for _, b := range []byte(S) {
		if b == ' ' {
			if beginningByte > 0 {
				sb.WriteByte(beginningByte)
			}
			sb.WriteString("ma")
			sb.WriteString(strings.Repeat("a", cnt))
			sb.WriteByte(' ')
			beginningByte = 0
			beginning = true
			cnt++
			continue
		}

		// 如果是单词的开始字符
		if beginning {
			// 如果是元音，写入
			if d[b] {
				sb.WriteByte(b)
			} else {
				beginningByte = b
			}
			beginning = false
		} else {
			sb.WriteByte(b)
		}
	}

	if beginningByte > 0 {
		sb.WriteByte(beginningByte)
	}
	sb.WriteString("ma")
	sb.WriteString(strings.Repeat("a", cnt))

	return sb.String()
}
