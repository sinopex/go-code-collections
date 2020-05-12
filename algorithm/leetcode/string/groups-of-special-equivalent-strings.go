// 893 特殊等价字符串组
// https://leetcode-cn.com/problems/groups-of-special-equivalent-strings/
package string

import (
	"bytes"
	"sort"
)

func numSpecialEquivGroups(A []string) int {
	var data = make(map[string]struct{})

	for _, str := range A {
		var xb, yb bytes.Buffer
		for i := 0; i < len(str); i++ {
			if i%2 == 0 {
				xb.WriteByte(str[i])
			} else {
				yb.WriteByte(str[i])
			}
			sort.Sort(Bytes(xb.Bytes()))
			sort.Sort(Bytes(yb.Bytes()))
		}
		data[xb.String()+yb.String()] = struct{}{}
	}

	return len(data)
}

type Bytes []byte

func (b Bytes) Len() int           { return len(b) }
func (b Bytes) Less(i, j int) bool { return b[i] < b[j] }
func (b Bytes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
