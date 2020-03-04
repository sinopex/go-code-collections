// 基数排序（Radix Sort）是非比较型排序算法，它和计数排序、桶排序一样，利用了“桶”的概念。
// 基数排序不需要进行记录关键字间的比较，是一种借助多关键字排序的思想对单逻辑关键字进行排序的方法。
// 比如数字100，它的个位、十位、百位就是不同的关键字。
// 参考链接：https://www.cnblogs.com/fengyumeng/p/10994279.html
//
// 最优时间复杂度：O(d(n+r))
// 平均时间复杂度：O(d(n+r))
// 最坏时间复杂度：O(d(n+r))
// 最坏空间复杂度：O(n+r)
// n:待排序列的个数
// r:桶的个数
// d:待排序的最高位数
package sorting

import (
	"math"
)

func Radix(arr []int) []int {
	sorted, length := copyArray(arr)

	// 获取最大值
	max := sorted[0]
	for i := 1; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// 10个存放位数的桶，默认都是空的
	var bucket [10][]int
	for i := 0; i < 10; i++ {
		bucket[i] = make([]int, 0)
	}

	location := 1
	for {
		dd := int(math.Pow10(location - 1))
		if max < dd {
			break
		}

		for i := 0; i < length; i++ {
			index := sorted[i] / dd % 10
			bucket[index] = append(bucket[index], sorted[i])
		}

		n := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < len(bucket[i]); j++ {
				sorted[n] = bucket[i][j]
				n++
			}
			// 将桶回收
			bucket[i] = make([]int, 0)
		}
		location++
	}

	return sorted
}
