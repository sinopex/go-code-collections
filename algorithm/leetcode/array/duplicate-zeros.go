// 1089. 复写零
// https://leetcode-cn.com/problems/duplicate-zeros/
package array

func duplicateZeros(arr []int) {
	moves := 0
	length := len(arr) - 1
	for left := 0; left <= length-moves; left++ {
		if arr[left] == 0 {
			if left == length-moves {
				arr[length] = 0
				length--
				break
			}
			moves++
		}
	}

	last := length - moves
	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+moves] = 0
			moves--
			arr[i+moves] = 0
		} else {
			arr[i+moves] = arr[i]
		}
	}
}

func duplicateZeros1(arr []int) {
	n := len(arr)
	t := make([]int, n)
	c := 0
	for i := 0; i < n; i++ {
		t[c] = arr[i]
		c++
		if c == n {
			break
		}
		if arr[i] == 0 {
			t[c] = arr[i]
			c++
		}
		if c == n {
			break
		}
	}

	copy(arr, t)
}
