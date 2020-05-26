// 1002. 查找常用字符
// https://leetcode-cn.com/problems/find-common-characters/
package array

func commonChars(A []string) []string {
	var table [26]int
	var result = make([]string, 0)
	for i, a := range A {
		var temp [26]int
		for _, c := range a {
			if i == 0 {
				table[c-'a']++
			} else {
				temp[c-'a']++
			}
		}
		if i != 0 {
			for j := 0; j < 26; j++ {
				if table[j] > 0 {
					if temp[j] < table[j] {
						table[j] = temp[j]
					}
				}
			}
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < table[i]; j++ {
			result = append(result, string(i+'a'))
		}
	}

	return result
}
