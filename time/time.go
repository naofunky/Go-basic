package main

import (
	"fmt"
	"time"
)

type MyTime struct {
	t time.Time
}

func main() {
	// // ローカル時間を取得
	// localTime := time.Now()

	// // UTCに変換
	// utcTime := localTime.UTC()

	// // ローカル時間とUTC時間を出力
	// fmt.Printf("Local time: %s\n", localTime)
	// fmt.Printf("UTC time: %s\n", utcTime)

	t1 := MyTime{t: time.Now()}
	t2 := MyTime{t: time.Now()}

	// invalid operation: t1 == t2 (struct containing time.Time cannot be compared)というコンパイルエラーになるのかなと思ったら処理が実行されてしまう。
	if t1.t.Equal(t2.t) {
		fmt.Println("t1 is equal to t2")
	} else {
		fmt.Println("t1 is not equal to t2")
	}

	// fmt.Printf("Local time: %s\n", t1)
	// fmt.Printf("Local time: %s\n", t2)
}
