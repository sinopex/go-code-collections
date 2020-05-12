// 1417. 重新格式化字符串
// https://leetcode-cn.com/problems/reformat-the-string/
package string

import (
	"bytes"
	"strings"
)

func reformat(s string) string {
	var digitBuf, letterBuf bytes.Buffer
	for _, x := range []byte(s) {
		if x >= '0' && x <= '9' {
			digitBuf.WriteByte(x)
		} else {
			letterBuf.WriteByte(x)
		}
	}
	digitNum, letterNum := digitBuf.Len(), letterBuf.Len()
	if letterNum > digitNum {
		digitBuf, letterBuf = letterBuf, digitBuf
		digitNum, letterNum = letterNum, digitNum
	}
	if digitNum-letterNum > 1 {
		return ""
	}

	var sb strings.Builder
	for {
		b1, err := digitBuf.ReadByte()
		if err != nil {
			break
		}
		sb.WriteByte(b1)

		b2, err := letterBuf.ReadByte()
		if err != nil {
			break
		}

		sb.WriteByte(b2)
	}

	return sb.String()
}
