// 1287. 有序数组中出现次数超过25%的元素
// https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array/
package array

func findSpecialInteger(arr []int) int {
	k := len(arr) / 4
	v, n := arr[0], 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == v {
			n++
			if n > k {
				break
			}
		} else {
			v = arr[i]
			n = 1
		}
	}
	return v
}
