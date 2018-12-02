package main

import (
	"fmt"
)

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	c <- 100

	// main()のfor-rangeに終了したことを通知する
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)

	// 並行実行
	go goroutine1(s, c)

	// チャネルに格納される(c <- xxx)値を随時取得
	//  close(c)が呼ばれると終了
	for i := range c {
		fmt.Println(i)
	}
}
