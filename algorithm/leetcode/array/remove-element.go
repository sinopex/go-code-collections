// 27. 移除元素
// https://leetcode-cn.com/problems/remove-element/
package array

// 直白易懂
// 找到不是目标值的位置，按顺序写入到左边
func removeElement(nums []int, val int) int {
	x, y := 0, 0

	for y < len(nums) {
		if nums[y] != val {
			nums[x] = nums[y]
			x++
		}
		y++
	}

	return x
}

// 前后双指针
func removeElement2(nums []int, val int) int {
	x, y := 0, len(nums)-1
	for x <= y {
		if nums[x] == val && nums[y] != val {
			nums[x], nums[y] = nums[y], nums[x]
			x++
			y--
		}

		if nums[x] != val {
			x++
		}

		if nums[y] == val {
			y--
		}
	}

	return x
}

// 顺向思维
func removeElement1(nums []int, val int) int {
	n := len(nums)

	var i, cursor int
	for i = 0; i < n; i++ {
		if cursor < i {
			cursor = i
		}
		if nums[i] != val {
			continue
		}
		for cursor += 1; cursor < n; cursor++ {
			if nums[cursor] != val {
				break
			}
		}
		if cursor == n {
			break
		}

		nums[i], nums[cursor] = nums[cursor], nums[i]
	}

	return i
}
