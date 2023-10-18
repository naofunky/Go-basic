package main

import (
	"fmt"
	"time"
)

func main() {
	// ローカル時間を取得
	localTime := time.Now()

	// UTCに変換
	utcTime := localTime.UTC()

	// ローカル時間とUTC時間を出力
	fmt.Printf("Local time: %s\n", localTime)
	fmt.Printf("UTC time: %s\n", utcTime)
}
