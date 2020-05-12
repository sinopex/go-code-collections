// 387. 字符串中的第一个唯一字符
// https://leetcode-cn.com/problems/first-unique-character-in-a-string/
package string

func firstUniqChar(s string) int {
	var letters [26]int

	for _, x := range []byte(s) {
		letters[x-'a']++
	}

	for i, x := range []byte(s) {
		if letters[x-'a'] == 1 {
			return i
		}
	}

	return -1
}

// func firstUniqChar(s string) int {
// 	var cache = make(map[byte][]int)
//
// 	for index, r := range []byte(s) {
// 		if _, ok := cache[r]; ok {
// 			cache[r][1]++
// 		} else {
// 			cache[r] = []int{index, 1}
// 		}
// 	}
//
// 	min := -1
// 	for _, v := range cache {
// 		if v[1] == 1 {
// 			if min == -1 {
// 				min = v[0]
// 			} else if min > v[0] {
// 				min = v[0]
// 			}
// 		}
// 	}
// 	return min
// }
