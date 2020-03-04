// 选择排序
// 选择排序（Selection Sort）是一种简单直观的排序算法。
// 它的基本思想就是，每一趟 n-i+1(i=1,2,...,n-1) 个记录中,
// 选取关键字最小的记录作为有序序列的第 i 个记录。
//
// 简单选择排序：
//
// 1. 在未排序序列中找到最小（大）元素，存放到排序序列的起始位置;
// 2. 在剩余未排序元素中继续寻找最小（大）元素，放到已排序序列的末尾;
// 3. 重复步骤 2，直到所有元素排序完毕
//
// 最优时间复杂度：O(n*n)
// 平均时间复杂度：O(n*n)
// 最坏时间复杂度：O(n*n)
// 最坏空间复杂度：O(1)

package sorting

func Selection(arr []int) []int {
	sorted, length := copyArray(arr)

	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if sorted[min] > sorted[j] {
				min = j
			}
		}
		sorted[i], sorted[min] = sorted[min], sorted[i]
	}

	return sorted
}
