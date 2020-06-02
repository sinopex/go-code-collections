// 122. 买卖股票的最佳时机 II
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
package array

func maxProfit(prices []int) int {
	num := len(prices)
	if num <= 1 {
		return 0
	}
	profit := 0
	index := 0
	for index < num-1 {
		// 涨价
		if prices[index] < prices[index+1] {
			profit += prices[index+1] - prices[index]
		}
		index += 1
	}
	return profit
}
