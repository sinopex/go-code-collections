// 38. 外观数列
// https://leetcode-cn.com/problems/count-and-say/
package leetcode

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	ans := "11"

	for i := 2; i < n; i++ {
		lastByte := ans[0]
		var buf bytes.Buffer
		count, s := 1, len(ans)
		for j := 1; j < s; j++ {
			// 如果探针遇到相同的值，增加同值计数器，并继续往前扫描
			if ans[j] == lastByte {
				count++
				if j == s-1 {
					buf.WriteString(strconv.Itoa(count))
					buf.WriteByte(lastByte)
				}
			} else {
				// 如果探针与base指针的字符不一样，
				// 则需要写入base字符
				buf.WriteString(strconv.Itoa(count))
				buf.WriteByte(lastByte)
				lastByte = ans[j] // 设置新的base指针
				count = 1         // 计数器清零

				// 如果到底了，则加入最后一个不同值的字符
				if j == s-1 {
					buf.WriteByte('1')
					buf.WriteByte(ans[j])
				}
			}
		}
		ans = buf.String()
	}

	return ans
}
