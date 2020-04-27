// 1108 IP地址无效化
package leetcode

import (
	"strings"
)

func defangIPaddr(address string) string {
	var sb strings.Builder
	for _, x := range address {
		if x == '.' {
			sb.WriteString("[.]")
		} else {
			sb.WriteString(string(x))
		}
	}

	return sb.String()
}
