// 888. 公平的糖果交换
// https://leetcode-cn.com/problems/fair-candy-swap/
package array

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB, tableA := 0, 0, make(map[int]struct{})
	for _, v := range A {
		sumA += v
		tableA[v] = struct{}{}
	}

	for _, v := range B {
		sumB += v
	}

	diff := (sumB - sumA) / 2
	for _, v := range B {
		if _, ok := tableA[v-diff]; ok {
			return []int{v - diff, v}
		}
	}
	return nil
}
