// 985. 查询后的偶数和
// https://leetcode-cn.com/problems/sum-of-even-numbers-after-queries/
package array

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	result := make([]int, len(A))
	sum := 0
	for _, v := range A {
		if v&1 == 0 {
			sum += v
		}
	}

	for k, _ := range A {
		value := queries[k][0]
		index := queries[k][1]

		oldValue := A[index]
		value += oldValue
		//
		if oldValue&1 == 0 && value&1 == 0 {
			sum = sum - oldValue + value
		} else if oldValue&1 == 0 {
			sum = sum - oldValue
		} else if value&1 == 0 {
			sum = sum + value
		} else {

		}

		result[k] = sum
		A[index] = value
	}

	return result
}
