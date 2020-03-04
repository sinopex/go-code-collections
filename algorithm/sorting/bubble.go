// 冒泡排序
// 最优时间复杂度：O(n)
// 平均时间复杂度：O(n*n)
// 最坏时间复杂度：O(n*n)
// 最坏空间复杂度：O(1)
package sorting

// 外层循环每遍历一次，将最小的数放在最前面
func Bubble(arr []int) []int {
	sorted, length := copyArray(arr)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	return sorted
}

// 改良版的，最好O(n)
// 将大数往后排放，每外层循环一次，搞定一个最大数
// 如果一次循环没有发生Swap，则前面的数字已经是正序的
// 无需在遍历，排序已经完成
func BubbleN(arr []int) []int {
	sorted, length := copyArray(arr)
	for i := 0; i < length-1; i++ {
		didSwap := false
		for j := 0; j < length-1-i; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				didSwap = true
			}
		}
		if didSwap == false {
			return sorted
		}
	}

	return sorted
}
