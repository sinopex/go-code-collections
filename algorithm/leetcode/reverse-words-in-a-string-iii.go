// 557 反转字符串中的单词III
// https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
package leetcode

import (
	"bytes"
)

func reverseWords(s string) string {
	var buf bytes.Buffer

	// 双指针，x保留单词的起点，y一直往前探路，碰到空格停止
	x, y, n := 0, 0, len(s)

	for y < n {
		// 发现空格了，需要将[x,y]之间的单词，反序压入输出缓冲区
		if s[y] == ' ' {
			for i := y - 1; i >= x; i-- {
				buf.WriteByte(s[i])
			}
			buf.WriteByte(' ')
			// y继续向前走，x和y在新的同一起点
			y++
			x = y
		} else {
			y++
		}
	}
	for i := n - 1; i >= x; i-- {
		buf.WriteByte(s[i])
	}

	return buf.String()
}
