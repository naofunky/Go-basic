package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// 現在のローカル時間を取得
	localTime := time.Now()
	// UTC時間に変換した時間を取得
	utcTime := localTime.UTC()

	t.Run("ローカル時間が期待される範囲内であるための検証", func(t *testing.T) {
		if localTime.Before(time.Now().Add(-time.Second)) || localTime.After(time.Now().Add(time.Second)) {
			t.Errorf("Local time is out of expected range: %s", localTime)
		}
	})

	t.Run("UTC時刻が期待される範囲内であるかの検証", func(t *testing.T) {
		if utcTime.Before(time.Now().Add(-time.Second)) || utcTime.After(time.Now().Add(time.Second)) {
			t.Errorf("UTC time is out of expected range: %s", utcTime)
		}
	})
}
