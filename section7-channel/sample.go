package main

import (
	"fmt"
)

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	// チャネルへ結果セット
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	// チャネルへ結果セット
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}

	// 並列実行結果を取得するためのチャネル変数を生成
	c := make(chan int)

	// 並行実行
	go goroutine1(s, c)
	go goroutine2(s, c)

	// 受信
	//  cにsumが入るまでブロッキングされる
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
