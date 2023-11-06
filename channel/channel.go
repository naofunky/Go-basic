package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// チャネルのバッファなしの初期化
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		ch <- 10
		time.Sleep(500 * time.Millisecond)
	}()
	fmt.Println(<-ch)
	wg.Wait()
}
