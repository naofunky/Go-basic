package main

import "fmt"

// func add(start, end int) int {
// 	sum := 0
// 	for i := start; i <= end; i++ {
// 		sum += i
// 	}
// 	return sum
// }

func main() {
	// // チャネルのバッファなしの初期化
	// ch := make(chan int)

	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	// ゴルーチンリーク
	// ch1 := make(chan int)

	// go func() {
	// 	fmt.Println(<-ch1)
	// }()
	// ch1 <- 10
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())

	// 逐次と並行の比較
	// start := time.Now()

	// // 逐次処理
	// sum := add(1, 100000000)
	// fmt.Println("Sequential:", sum)

	// elapsed := time.Since(start)
	// fmt.Printf("Sequential Time: %d ms\n", elapsed.Microseconds())

	// start = time.Now()

	// // 並行処理
	// var wg sync.WaitGroup
	// wg.Add(2)

	// var sum1, sum2 int

	// go func() {
	// 	defer wg.Done()
	// 	sum1 = add(1, 50000000)
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	sum2 = add(50000001, 100000000)
	// }()

	// wg.Wait()

	// sum = sum1 + sum2
	// fmt.Println("Concurrency:", sum)

	// elapsed = time.Since(start)
	// fmt.Printf("Concurrency Time: %d ms\n", elapsed.Microseconds())

	// バッファあり
	ch2 := make(chan int, 1)
	ch2 <- 10
	fmt.Println(<-ch2)
}
