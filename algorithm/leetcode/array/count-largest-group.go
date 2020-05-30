// 1399. 统计最大组的数目
// https://leetcode-cn.com/problems/count-largest-group/
package array

func countLargestGroup(n int) int {
	var result = make(map[int]int)
	var answer int
	for i := 1; i <= n; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		result[sum]++
	}

	max := 0
	for _, v := range result {
		if v > max {
			max = v
			answer = 1
		} else if v == max {
			answer++
		}
	}

	return answer
}
