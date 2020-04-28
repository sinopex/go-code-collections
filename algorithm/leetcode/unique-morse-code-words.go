// 804 唯一摩尔斯密码词
// https://leetcode-cn.com/problems/unique-morse-code-words/
package leetcode

import (
	"strings"
)

func uniqueMorseRepresentations(words []string) int {
	var dict = [...]string{
		".-", "-...", "-.-.", "-..", ".", "..-.",
		"--.", "....", "..", ".---", "-.-", ".-..",
		"--", "-.", "---", ".--.", "--.-", ".-.",
		"...", "-", "..-", "...-", ".--", "-..-",
		"-.--", "--..",
	}

	var result = make(map[string]struct{})
	for _, word := range words {
		var sb = strings.Builder{}
		for _, r := range word {
			sb.WriteString(dict[r-97])
		}

		result[sb.String()] = struct{}{}
	}

	return len(result)
}
