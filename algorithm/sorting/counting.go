// 计数排序（Counting Sort）是一种非比较性质的排序算法，利用了桶的思想。
// 它的核心在于将输入的数据值转化为键存储在额外开辟的辅助空间中，
// 也就是说这个辅助空间的长度取决于待排序列中的数据范围。
//
// 最优时间复杂度：O(n+r)
// 平均时间复杂度：O(n+r)
// 最坏时间复杂度：O(n+r)
// 最坏空间复杂度：O(n+r)
package sorting

func Counting(arr []int) []int {
	sorted, length := copyArray(arr)

	min, max := sorted[0], sorted[0]
	for i := 0; i < length; i++ {
		if sorted[i] > max {
			max = sorted[i]
		}
		if sorted[i] < min {
			min = sorted[i]
		}
	}

	r := max - min + 1
	buckets := make([]int, r)

	for i := 0; i < length; i++ {
		buckets[sorted[i]-min] += 1
	}

	index := 0
	// n是第几个桶，m是桶里存了几份
	for n, m := range buckets {
		for i := 0; i < m; i++ {
			sorted[index] = min + n
			index++
		}
	}

	return sorted
}
