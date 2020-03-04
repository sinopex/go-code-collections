// bucket sort 桶排序
//
// 桶排序的基本思路是将数据根据计算规则来分组，并将数据依次分配到对应分组中。
// 分配时可能出现某分组里有多个数据，那么再进行分组内排序。
// 然后得到了有序分组，最后把它们依次取出来放到数组中即实现了整体排序。
// 此处分类即是桶，计算规则是 index = value * (n-1) /k,
// (n为数据个数即桶个数，k为使全部数据分布在k*[0,1]空间中的正整数，k可取数据最大值)。
// 桶本身等价于一个二维数组。
// 注意点：是在插入桶的时候还是在插入桶之后进行桶内排序呢？我觉得入桶后比较好，
// 因为入桶后可以从桶内整体出发进行桶内排序，而不用挨个比较
// 动画演示 https://www.cs.usfca.edu/~galles/visualization/BucketSort.html
//
// k为桶数
// 最优时间复杂度：O(n+k)
// 平均时间复杂度：O(n+k)
// 最坏时间复杂度：O(n*n)
// 最坏空间复杂度：O(n)
package sorting

func Bucket(arr2 []int) []int {
	// num为桶数
	sorted, num := copyArray(arr2)
	// k（数组最大值）
	max := getMaxInArr(sorted)
	// 二维切片
	buckets := make([][]int, num)

	// 分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = sorted[i] * (num - 1) / max // 分配桶index = value * (n-1) /k
		buckets[index] = append(buckets[index], sorted[i])
	}
	// 桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			buckets[i] = Quick(buckets[i])    // 桶内排序选用之前的快排
			copy(sorted[tmpPos:], buckets[i]) // 将第i桶的数字，放序放入目标数组
			tmpPos += bucketLen
		}
	}

	return sorted
}

// 获取数组最大值
func getMaxInArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
