// 1309 解码字母到整数映射
// https://leetcode-cn.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
package string

import (
	"bytes"
)

func freqAlphabets(s string) string {
	var buf bytes.Buffer
	var dict = map[string]byte{
		"1":  'a',
		"2":  'b',
		"3":  'c',
		"4":  'd',
		"5":  'e',
		"6":  'f',
		"7":  'g',
		"8":  'h',
		"9":  'i',
		"10": 'j',
		"11": 'k',
		"12": 'l',
		"13": 'm',
		"14": 'n',
		"15": 'o',
		"16": 'p',
		"17": 'q',
		"18": 'r',
		"19": 's',
		"20": 't',
		"21": 'u',
		"22": 'v',
		"23": 'w',
		"24": 'x',
		"25": 'y',
		"26": 'z',
	}

	i, n := 0, len(s)
	for i < n {
		if i+2 < n && s[i+2] == '#' {
			buf.WriteByte(dict[s[i:i+2]])
			i += 3
		} else {
			buf.WriteByte(dict[s[i:i+1]])
			i++
		}
	}

	return buf.String()
}
