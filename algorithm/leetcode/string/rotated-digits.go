// 788. 旋转数字
// https://leetcode-cn.com/problems/rotated-digits/
package string

func rotatedDigits(N int) int {
	d := []int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}
	count := 0
	// 如果N>10，需要重新初始化数组
	if N >= 10 {
		d = append(d, make([]int, N-9)...)
	}
	for i := 0; i <= N; i++ {
		// 如果末位或前缀有坏数，则当前为坏数
		if d[i/10] == -1 || d[i%10] == -1 {
			d[i] = -1
			continue
		}
		// 如果前缀有好数，或者末位是好数，则一定为好数
		if d[i] = d[i/10] | d[i%10]; d[i] == 1 {
			count++
		}
	}

	return count
}
func rotatedDigits1(N int) int {
	g := 0
	bad := map[int]bool{3: true, 4: true, 7: true}
	good := map[int]bool{2: true, 5: true, 6: true, 9: true}
	isGoodsNumber := func(number int) bool {
		var hasGoodNumber bool
		for number > 0 {
			if number >= 1000 && number <= 9999 {
				x := number / 1000
				if bad[x] {
					return false
				}
				if good[x] {
					hasGoodNumber = true
				}
				number = number % 1000
			} else if number >= 100 && number <= 999 {
				x := number / 100
				if bad[x] {
					return false
				}
				if good[x] {
					hasGoodNumber = true
				}
				number = number % 100
			} else if number >= 10 && number <= 99 {
				x := number / 10
				if bad[x] {
					return false
				}
				if good[x] {
					hasGoodNumber = true
				}
				number = number % 10
			} else {
				if bad[number] {
					return false
				}
				if good[number] {
					hasGoodNumber = true
				}
				break
			}
		}

		return hasGoodNumber
	}

	for i := 2; i <= N; i++ {
		if isGoodsNumber(i) {
			g++
		}
	}

	return g
}
