// 415. 字符串相加
// https://leetcode-cn.com/problems/add-strings/
package string

func addStrings(num1 string, num2 string) string {
	n1, n2 := len(num1)-1, len(num2)-1
	var result []byte
	var flag, sum byte
	if n1 > n2 {
		result = make([]byte, 0, n1+1)
	} else {
		result = make([]byte, 0, n2+1)
	}

	for n1 >= 0 || n2 >= 0 {
		sum = 0
		if n1 >= 0 {
			sum += num1[n1] - '0'
			n1--
		}

		if n2 >= 0 {
			sum += num2[n2] - '0'
			n2--
		}

		sum += flag
		if sum >= 10 {
			sum %= 10
			flag = 1
		} else {
			flag = 0
		}
		result = append(result, sum+'0')
	}

	if flag == 1 {
		result = append(result, '1')
	}
	n := len(result)
	for x, y := 0, n-1; x < y; x, y = x+1, y-1 {
		result[x], result[y] = result[y], result[x]
	}
	return string(result)
}
