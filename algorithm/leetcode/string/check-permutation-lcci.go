// 面试题 判断是否互为字符重排
// https://leetcode-cn.com/problems/check-permutation-lcci/
package string

func CheckPermutation(s1 string, s2 string) bool {
	var a1, a2 [128]int
	for _, r := range s1 {
		a1[r-'a']++
	}
	for _, r := range s2 {
		a2[r-'a']++
	}

	return a1 == a2
}
