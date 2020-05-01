// 1170. 比较字符串最小字母出现频次
// https://leetcode-cn.com/problems/compare-strings-by-frequency-of-the-smallest-character/
package leetcode

func numSmallerByFrequency(queries []string, words []string) []int {
	wordsFnNum := [11]int{}
	for _, word := range words {
		wordsFnNum[_numSmallerByFrequency(word)]++
	}
	var result = make([]int, len(queries), len(queries))
	for index, query := range queries {
		n := _numSmallerByFrequency(query)
		for cnt, val := range wordsFnNum {
			if n < cnt {
				result[index] += val
			}
		}
	}

	return result
}

// 找到每个字符串的最小char的频次
func _numSmallerByFrequency(str string) int {
	var dict = [26]int{}
	for _, r := range str {
		dict[r-'a'] += 1
	}

	for _, v := range dict {
		if v > 0 {
			return v
		}
	}
	return 0
}
