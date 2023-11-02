package main

import (
	"log"
	"os"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// // 関数をゴルーチンとして実行
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("goroutine invoked")
	// }()
	// wg.Wait()
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	// fmt.Println("main func finish")

	// トレーサーの作成
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()
}
