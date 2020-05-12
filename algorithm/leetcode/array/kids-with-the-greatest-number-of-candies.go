// 1431. 拥有最多糖果的孩子
// https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
package array

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	if n == 0 {
		return nil
	}

	max := candies[0]
	for i := 1; i < n; i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	result := make([]bool, n)
	for i := 0; i < n; i++ {
		if candies[i]+extraCandies >= max {
			result[i] = true
		}
	}

	return result
}
