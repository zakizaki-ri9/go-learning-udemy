package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// 並行実行数
//  NewWeightedに指定する数分、並行実行が可能
var s *semaphore.Weighted = semaphore.NewWeighted(2)

func longProcess(ctx context.Context) {
	// // ロック取得有無を判定、取得できなかったらそのまま終了
	// isAcquire := s.TryAcquire(1)
	// if !isAcquire {
	// 	fmt.Println("Coul not get lock")
	// 	return
	// }
	// ロックを取得、失敗してもそのまま残るs
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	time.Sleep(1 * time.Second)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
