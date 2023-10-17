package main

import "testing"

func TestRune(t *testing.T) {

	t.Run("半角英数字の文字数カウント検証", func(t *testing.T) {
		message := "A"
		if len(message) != 1 {
			t.Errorf("length is not 1")
		}
	})

	t.Run("日本語の文字数カウント検証", func(t *testing.T) {
		message := "こんにちは"
		messageRune := []rune(message)
		if len(messageRune) != 5 {
			t.Errorf("character length is not 5")
		}
	})
}
