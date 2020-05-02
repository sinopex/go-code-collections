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

	// 先拿到n-1的数据
	preResult := countAndSay(n - 1)
	var buf bytes.Buffer
	lastByte, lastCount, size := preResult[0], 1, len(preResult)
	for j := 1; j < size; j++ {
		// 如果探针遇到相同的值，增加同值计数器，并继续往前扫描
		if preResult[j] == lastByte {
			lastCount++
		} else {
			// 如果探针与base指针的字符不一样，
			// 则需要写入base字符
			buf.WriteString(strconv.Itoa(lastCount))
			buf.WriteByte(lastByte)
			lastByte = preResult[j] // 设置新的base指针
			lastCount = 1           // 计数器清零
		}
	}

	buf.WriteString(strconv.Itoa(lastCount))
	buf.WriteByte(lastByte)

	return buf.String()
}
