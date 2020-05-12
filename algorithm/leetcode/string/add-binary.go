// 67. 二进制求和
// https://leetcode-cn.com/problems/add-binary/
package string

func addBinary(a string, b string) string {
	var buf []byte
	var flag int
	i, j := len(a)-1, len(b)-1
	if i > j {
		buf = make([]byte, 0, i+1)
	} else {
		buf = make([]byte, 0, j+1)
	}
	for i >= 0 || j >= 0 {
		t1, t2 := 0, 0
		if i >= 0 {
			t1 = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			t2 = int(b[j] - '0')
			j--
		}

		switch t1 + t2 + flag {
		case 3:
			flag = 1
			buf = append(buf, '1')
		case 2:
			flag = 1
			buf = append(buf, '0')
		case 1:
			flag = 0
			buf = append(buf, '1')
		case 0:
			flag = 0
			buf = append(buf, '0')
		}
	}
	if flag == 1 {
		buf = append(buf, '1')
	}
	for x, y := 0, len(buf)-1; x < y; x, y = x+1, y-1 {
		buf[x], buf[y] = buf[y], buf[x]
	}

	return string(buf)
}
