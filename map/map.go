package main

import "fmt"

func main() {
	// 宣言法
	m := map[string]string{}
	var nilmap map[string]int
	fmt.Println(len(nilmap))
	// fmt.Println(cap(nilmap))

	// 初期化
	// m := map[string]string{
	// 	"name": "Bob",
	// 	"age":  "18",
	// 	"sex":  "woman",
	// }

	// 値の追加する方法
	m["name"] = "Bob"
	m["age"] = "18"
	m["sex"] = "woman"

	// 値の取得はキーを指定する
	fmt.Println(m["name"])

	// 要素を取得する際に第二戻り値を用意すると、keyの有無を調べることができる
	// v, ok := m["sex"]
	// fmt.Println(v)
	// fmt.Println(ok)

	// for k, v := range m {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }

	// mapの削除機能は第一引数には削除対象のmap第二引数には削除する要素のキーを指定する
	// delete(m, "name")
	// fmt.Println(m)

	// fmt.Println(m)

	// 重複排除
	duplicatedArray := []int{1, 2, 3, 3, 4, 5}
	notDuplicatedArray := []int{}

	// 空の構造体を用いることで
	myMap := map[int]struct{}{}

	for _, v := range duplicatedArray {
		_, loopValue := myMap[v]
		// fmt.Println(v)
		// fmt.Println(loopValue)

		if !loopValue {
			myMap[v] = struct{}{}
			notDuplicatedArray = append(notDuplicatedArray, v)
		} else {
			fmt.Println("一致したキーがあるよ")
		}
	}

	fmt.Println(notDuplicatedArray)
}
