package main

import (
	"fmt"

	"github.com/zakizaki-ri9/go-learning/section8/mylib"
	"github.com/zakizaki-ri9/go-learning/section8/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Ave(s))

	mylib.Say()
	under.Hello()
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)

	// 大文字始まりでないと参照不可
	fmt.Println(mylib.Public)
	// fmt.Println(mylib.private)
}
