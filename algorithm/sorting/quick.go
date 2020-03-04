// 快速排序（QuickRefined Sort），是冒泡排序的改进版，之所以“快速”，是因为使用了分治法。
// 它也属于交换排序，通过元素之间的位置交换来达到排序的目的。
//
// 最优时间复杂度：O(n*log(n))
// 平均时间复杂度：O(n*log(n))
// 最坏时间复杂度：O(n*n)
// 最坏空间复杂度：O(log(n))
package sorting

func Quick(arr []int) []int {
	sorted, length := copyArray(arr)
	return quick(sorted, 0, length-1)
}

func quick(arr []int, _left, _right int) []int {
	left, right := _left, _right
	if left >= right {
		return arr
	}

	base := arr[left] // 取第一个元素作为基准元素
	for left != right {
		//从右边往左边扫描，找到第一个比基准元素小的元素
		for right > left && arr[right] >= base { // 这里要用>=，不然会出错
			right--
		}
		arr[left] = arr[right] //找到这种元素arr[right]后，与arr[left]交换

		// 从左边往右边扫描，找到第一个比基准元素大的元素
		for left < right && arr[left] <= base { // 这里要用<=，不然会出错
			left++
		}
		arr[right] = arr[left] //找到这中元素arr[left]后，与arr[right]交换
	}
	arr[right] = base
	quick(arr, _left, left-1)
	quick(arr, right+1, _right)
	return arr
}

// 在排序过程中，对基准的移动其实是多余的，因为只有一趟排序结束时，也就是 i = j 的位置才是基准的最终位置。
func QuickRefined(arr []int) []int {
	sorted, length := copyArray(arr)
	return quickRefined(sorted, 0, length-1)
}
func quickRefined(arr []int, left, right int) []int {
	start, end := left, right
	if start >= end {
		return arr
	}

	base := arr[left] // 取第一个元素作为基准元素
	for start != end {
		//从右边往左边扫描，找到第一个比基准元素小的元素
		for start < end && arr[end] >= base { // 这里要用>=，不然会出错
			end--
		}

		// 从左边往右边扫描，找到第一个比基准元素大的元素
		for start < end && arr[start] <= base { // 这里要用<=，不然会出错
			start++
		}

		// 将左边比基准信息大的，和右边比基准信息小的，互换
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		}
	}
	//将基准元素换到最终的位置
	arr[left], arr[start] = arr[start], base

	quickRefined(arr, left, start-1)
	quickRefined(arr, start+1, right)
	return arr
}
