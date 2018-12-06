/*
mylib is my kuso lib.
*/
package mylib

// 平均値を返却する
func Ave(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
