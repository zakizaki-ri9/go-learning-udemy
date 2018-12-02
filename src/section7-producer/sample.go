package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i*1000)
		wg.Done()
	}

	// Sleepしないと表示されない
	//  main()が終了してしまうため
	fmt.Println("################")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// chに順次格納される値を取得＆出力
	//  waitgroupが全て終了するまで待つ
	go consumer(ch, &wg)
	wg.Wait()

	// for-rangeを終了させるためのclose
	close(ch)

	// consumerの"#"を出力するためのsleep
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
