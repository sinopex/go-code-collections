// 929. 独特的电子邮件地址
// https://leetcode-cn.com/problems/unique-email-addresses/
package string

import (
	"bytes"
)

func numUniqueEmails(emails []string) int {
	var distinct = make(map[string]struct{})

	for _, email := range emails {
		var buf bytes.Buffer
		var isLock bool
		for i := 0; i < len(email); i++ {

			if email[i] == '.' {
				continue
			}

			if email[i] == '+' {
				isLock = true
			}

			if email[i] == '@' {
				buf.WriteString(string(email[i:]))
				break
			}

			if !isLock {
				buf.WriteByte(email[i])
			}

		}
		distinct[buf.String()] = struct{}{}
	}

	return len(distinct)
}
