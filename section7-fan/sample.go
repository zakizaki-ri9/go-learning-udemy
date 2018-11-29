package main

import (
	"fmt"
)

func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

// 関数変数名にfan-in fan-out を明記しておくとわかりやすい
func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		// iには0-9の値が順次格納される
		second <- i * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		// iには2の倍数(0-18 = {0-9} * 2)が順次格納される
		third <- i * 4
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(first, second)
	go multi4(second, third)
	for result := range third {
		// 8の倍数({0-9} * 2 * 4)の値を算出＆表示する
		fmt.Println(result)
	}
}
