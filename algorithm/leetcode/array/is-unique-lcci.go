// 面试题 01.01. 判定字符是否唯一
// https://leetcode-cn.com/problems/is-unique-lcci/
package array

func isUnique(astr string) bool {
	var mask uint32

	for _, b := range []byte(astr) {
		if mask&(1<<(b-'a')) > 0 {
			return false
		}
		mask |= 1 << (b - 'a')
	}

	return true
}
