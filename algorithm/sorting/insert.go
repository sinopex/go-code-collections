// 插入排序
// 直接插入排序（Straight Insertion Sort），是一种简单直观的排序算法，
// 它的基本操作是不断地将尚未排好序的数插入到已经排好序的部分，
// 好比打扑克牌时一张张抓牌的动作。
// 在冒泡排序中，经过每一轮的排序处理后，序列后端的数是排好序的；
// 而对于插入排序来说，经过每一轮的排序处理后，序列前端的数都是排好序的。
//
// 最优时间复杂度：O(n)
// 平均时间复杂度：O(n^2)
// 最坏时间复杂度：O(n^2)
// 最坏空间复杂度：O(1)
package sorting

func Insert(arr []int) []int {
	sorted, length := copyArray(arr)
	var i, j, temp int
	for i = 1; i < length; i++ {
		temp = sorted[i]
		for j = i; j > 0 && sorted[j-1] > temp; j-- {
			sorted[j] = sorted[j-1]
		}
		sorted[j] = temp
	}
	return sorted
}
