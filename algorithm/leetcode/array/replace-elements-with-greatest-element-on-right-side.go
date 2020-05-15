// 1299. 将每个元素替换为右侧最大元素
// https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/
package array

func replaceElements(arr []int) []int {
	max := -1
	var temp int
	for n := len(arr) - 1; n >= 0; n-- {
		temp = arr[n]
		arr[n] = max
		if temp > max {
			max = temp
		}
	}

	return arr
}
