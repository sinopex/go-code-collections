// 希尔排序
// 希尔排序是第一个突破 O(n²) 的排序算法，是直接插入排序的改进版，
// 又称“缩小增量排序”（Diminishing Increment Sort）。
// 它与直接插入排序不同之处在于，它会优先比较距离较远的元素。
//
// 最优时间复杂度：O(n)
// 平均时间复杂度：O(n*log2(n))
// 最坏时间复杂度：O(n*n)
// 最坏空间复杂度：O(1)
//
// 参考链接：https://blog.csdn.net/weixin_41190227/article/details/86600821
package sorting

func Shell(arr []int) []int {
	sorted, length := copyArray(arr)
	gap := length / 2
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := sorted[i]
			prevIndex := i - gap
			for prevIndex >= 0 && sorted[prevIndex] > temp {
				sorted[prevIndex+gap] = sorted[prevIndex]
				prevIndex -= gap
			}
			sorted[prevIndex+gap] = temp
		}
		gap = gap / 2
	}
	return sorted
}
