package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)

	ch <- 100
	fmt.Println(len(ch))

	ch <- 200
	fmt.Println(len(ch))

	// for-rangeで取り出す場合は一旦クローズすること
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}

	// 取り出したらチャネルの数が減っていることを確認
	fmt.Println(len(ch))
}
