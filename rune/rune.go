package main

import "fmt"

func main() {

	message := "A"
	fmt.Println(len(message))

	// 日本語だとおかしなことになる
	message1 := "あ"
	// fmt.Println(len(message1))

	// runeのスライスに変換するとうまくカウントすることができる
	message2 := []rune(message1)
	fmt.Println(len(message2))
}
