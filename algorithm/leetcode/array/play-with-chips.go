// 1217. 玩筹码
// https://leetcode-cn.com/problems/play-with-chips/
package array

func minCostToMoveChips(chips []int) int {
	odd, even := 0, 0
	for _, v := range chips {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd > even {
		return even
	}
	return odd
}
