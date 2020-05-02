// 1071. 字符串的最大公因子
// https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/
package leetcode

import (
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	a, b := len(str1), len(str2)
	g := gcd(a, b)
	candidate := str1[0:g]
	if strings.Repeat(candidate, a/g) == str1 && strings.Repeat(candidate, b/g) == str2 {
		return candidate
	}
	return ""
}

func gcd(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < b {
		a, b = b, a
	}

	for b > 0 {
		a, b = b, a%b
	}

	return a
}
