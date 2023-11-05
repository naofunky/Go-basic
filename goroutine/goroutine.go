package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
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

	// 他の処理の後にファイルを閉じる関数をdeferで遅延発火
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	// tracerをtrace.outファイルを開くのにバリデーション
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}

	defer trace.Stop()
	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

	// 逐次処理
	// task(ctx, "task1")
	// task(ctx, "task2")
	// task(ctx, "task3")

	// ゴルーチン処理
	var wg sync.WaitGroup
	wg.Add(3)
	go cTack(ctx, &wg, "Task1")
	go cTack(ctx, &wg, "Task2")
	go cTack(ctx, &wg, "Task3")
	wg.Wait()
	fmt.Println("main func finish")
}

// 逐次関数処理
// func task(ctx context.Context, name string) {
// 	defer trace.StartRegion(ctx, name).End()
// 	time.Sleep(time.Second)
// 	fmt.Println(name)
// }

// goroutineを利用してタスクを並行処理
func cTack(ctx context.Context, wg *sync.WaitGroup, name string) {
	// deferのメソッドチェーンは後ろの関数のみ遅延
	defer trace.StartRegion(ctx, name).End()
	// ゴルーチンの完了をここで終える
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}
