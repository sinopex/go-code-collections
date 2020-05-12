// 859. 亲密字符串
// https://leetcode-cn.com/problems/buddy-strings/
package string

func buddyStrings(A string, B string) bool {
	na, nb := len(A), len(B)
	if na != nb {
		return false
	}

	if na < 2 {
		return false
	}

	// 如果A、B相等，则查询是否有字符超过2个
	if A == B {
		chars := [26]int{}
		for _, v := range A {
			chars[v-'a']++
			if chars[v-'a'] == 2 {
				return true
			}
		}
		return false
	}

	// 不相等，则需要置换位置
	prev := -1
	replaced := false
	for i := 0; i < na; i++ {
		if A[i] != B[i] {
			if replaced == true {
				return false
			}
			// 第一个位置
			if prev == -1 {
				prev = i
			} else {
				// 第二个位置，直接比较
				if A[prev] == B[i] && A[i] == B[prev] {
					replaced = true
				} else {
					return false
				}
			}
		}

	}

	return replaced
}
