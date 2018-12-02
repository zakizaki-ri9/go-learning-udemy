package main

import (
	"fmt"
	"time"
)

func main() {
	// チャネルを返してくれる
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

OuterLoop:
	for {
		select {
		case t := <-tick:
			fmt.Println("tick.", t)
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop // ラベル「OuterLoop」の部分から抜ける
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	fmt.Println("###############")
}
