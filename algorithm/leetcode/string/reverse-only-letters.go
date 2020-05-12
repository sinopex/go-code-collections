// 917. 仅仅反转字母
// https://leetcode-cn.com/problems/reverse-only-letters/
package string

import (
	"bytes"
)

func reverseOnlyLetters(S string) string {
	var buf bytes.Buffer
	var ret bytes.Buffer
	n := len(S)

	for i := n - 1; i >= 0; i-- {
		if isValidAlphabet(S[i]) {
			buf.WriteByte(S[i])
		}
	}

	for i := 0; i < n; i++ {
		if !isValidAlphabet(S[i]) {
			ret.WriteByte(S[i])
		} else {
			b, _ := buf.ReadByte()
			ret.WriteByte(b)
		}
	}

	return ret.String()
}

func isValidAlphabet(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}
