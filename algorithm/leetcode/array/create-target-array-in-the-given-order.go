// 1389. 按既定顺序创建目标数组
// https://leetcode-cn.com/problems/create-target-array-in-the-given-order/
package array

// 这是我写的
// BenchmarkCreateTargetArray-4      300073              4090 ns/op             848 B/op         12 allocs/op
func createTargetArray2(nums []int, index []int) []int {
	n := len(index)
	target := make([]int, n)
	for i, v := range index {
		target = append(target[0:v], append([]int{nums[i]}, target[v:]...)...)[0:n]
	}

	return target
}

// 这是网上大牛写的
// Go 圣经里讲的，插入应该怎么插，就这么插，纠结不能用函数的，当然也可以循环移动
// BenchmarkCreateTargetArray-4      727120              1541 ns/op              80 B/op          1 allocs/op
func createTargetArray(nums []int, index []int) []int {
	var res = make([]int, len(index))
	for k, v := range index {
		copy(res[v+1:], res[v:])
		res[v] = nums[k]
	}

	return res
}
