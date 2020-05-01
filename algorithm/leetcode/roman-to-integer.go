// 13. 罗马数字转整数
// https://leetcode-cn.com/problems/roman-to-integer/
package leetcode

func romanToInt(s string) int {
	var num int

	length := len(s)
	for i := 0; i < length; i++ {
		switch s[i] {
		case 'I':
			if i+1 < length && s[i+1] == 'V' {
				num += 4
				i++
			} else if i+1 < length && s[i+1] == 'X' {
				num += 9
				i++
			} else {
				num += 1
			}
		case 'V':
			num += 5
		case 'X':
			if i+1 < length && s[i+1] == 'L' {
				num += 40
				i++
			} else if i+1 < length && s[i+1] == 'C' {
				num += 90
				i++
			} else {
				num += 10
			}
		case 'L':
			num += 50
		case 'C':
			if i+1 < length && s[i+1] == 'D' {
				num += 400
				i++
			} else if i+1 < length && s[i+1] == 'M' {
				num += 900
				i++
			} else {
				num += 100
			}
		case 'D':
			num += 500
		case 'M':
			num += 1000
		}
	}

	return num
}

func romanToInt1(s string) int {
	var dict = map[string]int{
		"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
		"XL": 40, "L": 50, "XC": 90, "C": 100,
		"CD": 400, "D": 500, "CM": 900, "M": 1000,
	}

	var n int
	var buf = []byte(s)
	i, j := len(buf)-1, len(buf)-2
	for i >= 0 {
		if j >= 0 {
			if v, ok := dict[string(buf[j:i+1])]; ok {
				n += v
				i -= 2
				j -= 2
				continue
			}
		}
		n += dict[string(buf[i])]
		i--
		j--
	}

	return n
}
