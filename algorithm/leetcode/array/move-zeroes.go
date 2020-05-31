// 283. 移动零
// https://leetcode-cn.com/problems/move-zeroes/
package array

func moveZeroes(nums []int) {
	notZeroIndex := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] != 0 {
			continue
		}
		if notZeroIndex < i {
			notZeroIndex = i
		}

		for notZeroIndex < n {
			// 后排找到一个不是0的数值，与前排的0替换
			if nums[notZeroIndex] != 0 {
				nums[i], nums[notZeroIndex] = nums[notZeroIndex], nums[i]
				break
			}
			notZeroIndex++
		}

		// 后面没有非零数字了
		if notZeroIndex == n {
			break
		}
	}
}
