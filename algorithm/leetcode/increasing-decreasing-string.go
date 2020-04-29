// 1370 上升下降的字符串
// https://leetcode-cn.com/problems/increasing-decreasing-string/
package leetcode

import (
	"strings"
)

func sortString(s string) string {
	var buckets [26]int
	for _, r := range s {
		buckets[r-'a'] += 1
	}

	var sb strings.Builder
	var total int
	for {
		total = 0
		// 正序
		for i := 0; i < len(buckets); i++ {
			if buckets[i] > 0 {
				sb.WriteByte(byte(i + 97))
				buckets[i]--
				total++
			}
		}
		// 没了
		if total == 0 {
			goto finish
		}

		total = 0
		// 反序
		for i := len(buckets) - 1; i >= 0; i-- {
			if buckets[i] > 0 {
				sb.WriteByte(byte(i + 97))
				buckets[i]--
				total++
			}
		}
		if total == 0 {
			goto finish
		}
	}

finish:
	return sb.String()
}

// 暴力方法
// func sortString(s string) string {
// 	// 先排序
// 	w := strings.Split(s, "")
// 	sort.Strings(w)
// 	s = strings.Join(w, "")
//
// 	var sb strings.Builder
// 	buf := []byte(s)
// 	var temp bytes.Buffer
// 	var picked map[byte]bool
// 	for len(buf) > 0 {
// 		temp = bytes.Buffer{}
// 		picked = make(map[byte]bool)
// 		// 走一遍正序
// 		for i := 0; i < len(buf); i++ {
// 			if picked[buf[i]] {
// 				temp.WriteByte(buf[i])
// 			} else {
// 				picked[buf[i]] = true
// 				// 将正序的压入输出的缓冲区
// 				sb.WriteByte(buf[i])
// 			}
// 		}
//
// 		buf = temp.Bytes()
// 		temp = bytes.Buffer{}
// 		picked = make(map[byte]bool)
//
// 		// 走一遍反序
// 		var reverseByte = make([]byte, 0)
// 		for i := 0; i < len(buf); i++ {
// 			if picked[buf[i]] {
// 				temp.WriteByte(buf[i])
// 			} else {
// 				picked[buf[i]] = true
// 				reverseByte = append(reverseByte, buf[i])
// 			}
// 		}
// 		buf = temp.Bytes()
// 		// 将反序的压入输出的缓冲区
// 		for i := len(reverseByte) - 1; i >= 0; i-- {
// 			sb.WriteByte(reverseByte[i])
// 		}
// 	}
// 	return sb.String()
// }
