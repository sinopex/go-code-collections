// 459. 重复的子字符串
// https://leetcode-cn.com/problems/repeated-substring-pattern/
package string

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:len(s)*2-1], s)
}

// func repeatedSubstringPattern(s string) bool {
// 	n := len(s)
// 	for i := 1; i <= n/2; i++ {
// 		if n%i == 0 {
// 			if strings.Repeat(s[0:i], n/i) == s {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
