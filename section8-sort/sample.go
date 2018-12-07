package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{3, 9, 5, 1, 0}
	s := []string{"d", "p", "a"}
	p := []struct {
		Name string
		Age  int
	}{
		{"ZZZ", 30},
		{"OOO", 40},
		{"NNN", 10},
		{"AAA", 20},
	}

	// ソート前
	fmt.Println(i, s, p)

	// ソート実行
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(
		p,
		func(i, j int) bool {
			return p[i].Name < p[j].Name
		},
	)
	sort.Slice(
		p,
		func(i, j int) bool {
			return p[i].Age < p[j].Age
		},
	)

	// ソート後
	fmt.Println(i, s, p)
}
