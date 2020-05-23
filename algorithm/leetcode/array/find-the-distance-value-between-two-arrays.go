// 1385. 两个数组间的距离值
// https://leetcode-cn.com/problems/find-the-distance-value-between-two-arrays/
package array

import (
	"math"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var result int
	var flag bool
	for _, m := range arr1 {
		for _, n := range arr2 {
			if math.Abs(float64(m-n)) <= float64(d) {
				flag = true
				break
			}
		}
		if !flag {
			result++
		} else {
			flag = false
		}
	}

	return result
}
