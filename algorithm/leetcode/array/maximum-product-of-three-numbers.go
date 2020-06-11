// 628. 三个数的最大乘积
// https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
package array

import (
	"math"
	"sort"
)

func maximumProduct(nums []int) int {
	bottom1, bottom2 := math.MaxInt32, math.MaxInt32
	top1, top2, top3 := -math.MaxInt32, -math.MaxInt32, -math.MaxInt32

	for _, v := range nums {
		if v <= bottom1 {
			bottom1, bottom2 = v, bottom1
		} else if v <= bottom2 {
			bottom2 = v
		}

		if v >= top1 {
			top1, top2, top3 = v, top1, top2
		} else if v >= top2 {
			top2, top3 = v, top2
		} else if v >= top3 {
			top3 = v
		}
		// fmt.Printf("(%d,%d,%d,%d,%d)\n", bottom1, bottom2, top3, top2, top1)
	}
	return max(top1*top2*top3, top1*top2*bottom1, bottom1*bottom2*top1)
}

func maximumProduct1(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(
		nums[0]*nums[1]*nums[2],
		nums[0]*nums[1]*nums[n-1],
		nums[0]*nums[n-2]*nums[n-1],
		nums[n-3]*nums[n-2]*nums[n-1],
	)
}
func max(v ...int) int {
	r := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > r {
			r = v[i]
		}
	}
	return r
}
