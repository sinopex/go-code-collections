// 443. 压缩字符串
// https://leetcode-cn.com/problems/string-compression/
package string

import (
	"strconv"
)

func compress(chars []byte) int {
	n := len(chars)
	if n == 0 {
		return 0
	}
	writePosition := 0 // 记录写的位置
	// 从0开始读，每次读完一个字符(包括后续所有连续相同的字符)，从写的位置写入
	for readStart := 0; readStart < n; {
		readEnd, d := readStart+1, 1
		// 读这个字符的后面有多少个相同的字符
		for readEnd < n {
			if chars[readStart] != chars[readEnd] {
				break
			}
			readEnd++
			d++
		}

		// 写入当前字符
		chars[writePosition] = chars[readStart]
		// 写入当前字符的重复数字
		if d > 1 {
			for _, v := range []byte(strconv.Itoa(d)) {
				writePosition++
				chars[writePosition] = v
			}
		}
		// 当前字符写完之后，写入指针往后挪一步
		writePosition++
		readStart = readEnd
	}

	return writePosition
}
