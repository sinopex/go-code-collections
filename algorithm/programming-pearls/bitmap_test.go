package programming_pears

import (
	"fmt"
	"testing"
)

func TestBitMap_Has(t *testing.T) {
	bitmap := New()
	bitmap.Add(18)
	bitmap.Add(271)
	bitmap.Add(100000091)
	fmt.Println(bitmap)
	if !bitmap.Has(18) {
		t.Errorf("error")
	}
	if bitmap.Has(10000) {
		t.Errorf("error")
	}
}
