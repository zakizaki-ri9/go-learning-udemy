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
		{"ZZZ", 5},
		{"NNN", 10},
		{"AAA", 20},
		{"OOO", 30},
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

	// ソート後
	fmt.Println(i, s, p)
}
