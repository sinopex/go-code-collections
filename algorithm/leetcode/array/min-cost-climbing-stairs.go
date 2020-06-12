// 746. 使用最小花费爬楼梯
// https://leetcode-cn.com/problems/min-cost-climbing-stairs/
package array

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	a, b := 0, 0
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 2; i <= n; i++ {
		a, b = b, min(b+cost[i-1], a+cost[i-2])
	}
	return b
}
