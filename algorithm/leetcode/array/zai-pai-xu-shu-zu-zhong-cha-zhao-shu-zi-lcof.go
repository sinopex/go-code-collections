// 面试题53 - I. 在排序数组中查找数字 I
// https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
package array

func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	right := i
	if j >= 0 && nums[j] != target {
		return 0
	}

	i = 0
	for i <= j {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	left := j
	return right - left - 1
}

func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			count := 1
			// 往左蔓延
			for temp := middle - 1; temp >= left; temp-- {
				if nums[temp] != target {
					break
				}
				count++
			}

			// 往右蔓延
			for temp := middle + 1; temp <= right; temp++ {
				if nums[temp] != target {
					break
				}
				count++
			}
			return count
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return 0
}
