package programming_pears

import (
	"bytes"
	"fmt"
)

// CopyFrom: https://zhuanlan.zhihu.com/p/37285799

type BitMap struct {
	words  []uint64
	length int
}

func New() *BitMap {
	return &BitMap{}
}

func (bitmap *BitMap) Has(num int) bool {
	word, bit := num/64, uint(num%64)
	return word < len(bitmap.words) && bitmap.words[word]&(1<<bit) != 0
}

func (bitmap *BitMap) Add(num int) {
	// int8个字节，一个字节8位，共64位
	word, bit := num/64, uint(num%64)
	// 如果当前长度不够，则动态扩容
	// 如：当前是空的slice，num要放到第5位，则先逐步分配内存
	for word >= len(bitmap.words) {
		bitmap.words = append(bitmap.words, 0)
	}

	// 判断num是否已经存在bitmap中
	if bitmap.words[word]&(1<<bit) == 0 {
		bitmap.words[word] |= 1 << bit
		bitmap.length++
	}
}

func (bitmap *BitMap) Len() int {
	return bitmap.length
}

func (bitmap *BitMap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range bitmap.words {
		if v == 0 {
			continue
		}

		for j := uint(0); j < 64; j++ {
			if v&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				_, _ = fmt.Fprintf(&buf, "%d", 64*uint(i)+j)
			}
		}
	}
	buf.WriteByte('}')
	_, _ = fmt.Fprintf(&buf, "\nLength: %d", bitmap.length)
	return buf.String()
}
