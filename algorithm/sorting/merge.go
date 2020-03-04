// 归并排序（Merge Sort）是建立在归并操作上的一种排序算法。它和快速排序一样，采用了分治法。
//
// 归并的含义是将两个或两个以上的有序表组合成一个新的有序表。
// 也就是说，从几个数据段中逐个选出最小的元素移入新数据段的末尾，使之有序。
// 假如初始序列含有 n 个记录，则可以看成是 n 个有序的子序列，每个子序列的长度为 1。
// 然后两两归并，得到 n/2 个长度为2或1的有序子序列；再两两归并，……，
// 如此重复，直到得到一个长度为 n 的有序序列为止，这种排序方法称为 二路归并排序
//
// 最优时间复杂度：O(n*log(n))
// 平均时间复杂度：O(n*log(n))
// 最坏时间复杂度：O(n*log(n))
// 最坏空间复杂度：O(n)
package sorting

func Merge(arr []int) []int {
	sorted, length := copyArray(arr)
	divide(sorted, 0, length-1)
	return sorted
}

// 将arr[l..m]和 arr[m+1..r]合并
func merge(arr []int, l, r, m int) {
	leftSize, rightSize := m-l+1, r-m
	left, right := make([]int, leftSize), make([]int, rightSize)
	// 以M为分割线，将原数组分成左右子数组
	// 左边排序好的数据就位
	for i := l; i <= m; i++ {
		left[i-l] = arr[i]
	}
	// 右边排序好的数据就位
	for i := m + 1; i <= r; i++ {
		right[i-m-1] = arr[i]
	}

	// 再合成一个有序数组：从两个序列中选出最小值一次插入
	i, j, k := 0, 0, l
	for i < leftSize && j < rightSize {
		// 左边当前序列小，将左边的值压入最终的数组序号
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			// 右边当前序列小，将右边的值压入最终的数组序号
			arr[k] = right[j]
			j++
		}
		k++
	}

	// 如果左边没有拿完，则全部垫到后面
	for i < leftSize {
		arr[k] = left[i]
		i++
		k++
	}

	// 同理，处理右边
	for j < rightSize {
		arr[k] = right[j]
		j++
		k++
	}
}

func divide(arr []int, l, r int) {
	// 终止条件，只有一个元素
	if l == r {
		return
	}

	m := (l + r) / 2
	divide(arr, l, m)
	divide(arr, m+1, r)
	merge(arr, l, r, m)
}
