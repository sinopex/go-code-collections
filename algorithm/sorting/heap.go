// 堆排序（Heap Sort）是指利用堆这种数据结构所设计的一种排序算法。
//
// 堆的特点：
// 1. 一颗完全二叉树（也就是会所生成节点的顺序是：从上往下、从左往右）
// 2. 每一个节点必须满足父节点的值不大于/不小于子节点的值
//
// 实现堆排序需要解决两个问题：
// 1. 如何将一个无序序列构建成堆？
// 2. 如何在输出堆顶元素后，调整剩余元素成为一个新的堆？
//
// 以升序为例，算法实现的思路为：
// 1. 建立一个 build_heap 函数，将数组 tree[0,...n-1] 建立成堆，
//    n 表示数组长度。函数里需要维护的是所有节点的父节点，
//    最后一个子节点下标为 n-1，那么它对应的父节点下标就是 (n-1-1)/2。
// 2. 构建完一次堆后，最大元素就会被存放在根节点 tree[0]。
//    将 tree[0] 与最后一个元素交换，每一轮通过这种不断将最大元素后移的方式，来实现排序。
// 3. 而交换后新的根节点可能不满足堆的特点了，因此需要一个调整函数
//    heapify 来对剩余的数组元素进行最大堆性质的维护。如果 tree[i] 表示其中的某个节点，
//    那么 tree[2*i+1] 是左孩子，tree[2*i+2] 是右孩子，选出三者中的最大元素的下标，
//    存放于 max 值中，若 max 不等于 i，则将最大元素交换到 i 下标的位置。
//    但是，此时以 tree[max] 为根节点的子树可能不满足堆的性质，需要递归调用自身。
//
// 最优时间复杂度：O(n*log(n))
// 平均时间复杂度：O(n*log(n))
// 最坏时间复杂度：O(n*log(n))
// 最坏空间复杂度：O(log(n))
package sorting

func Heap(arr []int) []int {
	sorted, length := copyArray(arr)
	// 建堆
	buildHeap(sorted, length)

	// 每迭代一次，将堆顶的最大元素设置到数组的最后一位
	for i := length - 1; i >= 0; i-- {
		sorted[0], sorted[i] = sorted[i], sorted[0]
		// 重新建堆
		heapify(sorted, i, 0)
	}

	return sorted
}

// 将数组堆化
func buildHeap(tree []int, length int) {
	lastNode := length - 1
	parent := (lastNode - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(tree, length, i)
	}
}

// 对数组进行最大堆维护
// tree，n表示序列长度，i表示父节点下标
func heapify(tree []int, n, i int) {
	if i >= n {
		return
	}
	max, left, right := i, 2*i+1, 2*i+2
	if left < n && tree[left] > tree[max] {
		max = left
	}

	if right < n && tree[right] > tree[max] {
		max = right
	}

	// 说明有较大的值在子节点，需要跟父节点互换
	if max != i {
		tree[max], tree[i] = tree[i], tree[max]
		heapify(tree, n, max)
	}
}
