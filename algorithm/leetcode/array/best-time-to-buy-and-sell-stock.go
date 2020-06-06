// 121. 买卖股票的最佳时机
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
package array

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, max := prices[0], 0

	for i := 1; i < len(prices); i++ {
		val := prices[i] - min
		if val > max {
			max = val
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return max
}
