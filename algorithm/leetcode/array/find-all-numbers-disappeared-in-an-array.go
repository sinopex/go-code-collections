// 448. 找到所有数组中消失的数字
// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
package array

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)
	var abs = func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}

	for _, v := range nums {
		nums[abs(v)-1] = -abs(nums[abs(v)-1])
	}

	for k, v := range nums {
		if v > 0 {
			result = append(result, k+1)
		}
	}

	return result
}
