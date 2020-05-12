// 1313. 解压缩编码列表
// https://leetcode-cn.com/problems/decompress-run-length-encoded-list/
package array

func decompressRLElist(nums []int) []int {
	result := make([]int, 0)

	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			result = append(result, nums[i+1])
		}
	}

	return result
}
