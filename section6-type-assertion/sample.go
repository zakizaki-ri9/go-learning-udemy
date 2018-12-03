package main

import (
	"fmt"
)

func do(i interface{}) {
	ii := i.(int)
	// i = ii * 2
	ii *= 2
	fmt.Println(ii)
}

func do2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + " > ok!!")
	default:
		fmt.Printf("unknown!! > %T\n", v)
	}
}

func main() {
	var i interface{} = 10
	do(i)

	var s interface{} = "aaa"
	// コンパイルは通るが、タイプアサーション時(i.(int))にエラーが発生する
	// do(s)
	do2(s)

	do2(1.1)
}
