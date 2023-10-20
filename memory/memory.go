package main

import (
	"fmt"
	"runtime"
)

// 状態を出力する関数
func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// 現在のヒープの状態を出力
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	// ヒープに割り当てられた累計量
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	// システムがGoランタイムに割り当てたメモリ量
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	// そしてガベージコレクションの回数
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

// バイトをメガバイトに変換するヘルパー関数
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// メイン関数
func main() {
	printMemStats()

	// 思い処理なので注意
	data := make([]int, 1<<25) // 2^25要素のスライスを作成
	for i := range data {
		data[i] = i
	}
	printMemStats()
}
