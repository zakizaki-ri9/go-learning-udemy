package main

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	// ストックを1つ減らす
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup

	// 1つの並列処理があるということをストック
	wg.Add(1)

	go goroutine("world", &wg)
	normal("hello")

	//time.Sleep(1000 * time.Millisecond)
	// ストックがなくなるまで待つ
	//  0にならないと例外が発生する
	wg.Wait()
}
