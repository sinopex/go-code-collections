// 830. 较大分组的位置
// https://leetcode-cn.com/problems/positions-of-large-groups/
package array

func largeGroupPositions(S string) [][]int {
	result := make([][]int, 0)
	left := 0
	for left < len(S)-1 {
		var right int
		count := 1
		for right = left + 1; right < len(S); right++ {
			if S[left] == S[right] {
				count++
			} else {
				break
			}
		}
		if count >= 3 {
			result = append(result, []int{left, right - 1})
		}
		left = right
	}

	return result
}
