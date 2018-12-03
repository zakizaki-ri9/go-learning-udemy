package mylib

import "testing"

// 平均値を返す
func TestAve(t *testing.T) {
	v := Ave([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Mistake")
	}
}
