package main

import (
	"fmt"
	"time"
)

// type MyTime struct {
// 	t time.Time
// }

func main() {
	// // ローカル時間を取得
	localTime := time.Now()

	// UTCに変換
	utcTime := localTime.UTC()

	// ローカル時間とUTC時間を出力
	fmt.Printf("Local time: %s\n", localTime)
	fmt.Printf("UTC time: %s\n", utcTime)

	// 不等号の比較の検証
	// second := int((8 * time.Hour).Seconds())
	// location := time.FixedZone("Asia/Tokyo", second)

	// fmt.Println(second)
	// fmt.Println(location)

	// timeTokyo := time.Date(2023, 10, 19, 8, 0, 0, 0, location)
	// timeUTC := time.Date(2023, 10, 19, 0, 0, 0, 0, time.UTC)

	// fmt.Println(timeTokyo)
	// fmt.Println(timeUTC)

	// compareCharacterEqual := timeTokyo == timeUTC
	// compareEqual := timeTokyo.Equal(timeUTC)

	// fmt.Println(compareCharacterEqual)
	// fmt.Println(compareEqual)
}
