package main

import "fmt"

func main() {
	// 宣言法
	// m := map[string]string{}
	// var nilmap map[string]int
	// fmt.Println(len(nilmap))
	// fmt.Println(m)

	// 初期化
	mo := map[string]string{
		"name": "Bob",
		"age":  "18",
		"sex":  "woman",
	}

	// 値の追加する方法
	mo["name"] = "Bob"
	mo["age"] = "18"
	mo["sex"] = "woman"

	// 値の取得はキーを指定する
	fmt.Println(mo["name"])

	// 要素を取得する際に第二戻り値を用意すると、keyの有無を調べることができる
	v, ok := mo["sex"]
	fmt.Println("keyに対しての値確認:", v)
	fmt.Println("keyの存在確認:", ok)

	for k, v := range mo {
		fmt.Println(
			"key:", k, "value:", v,
		)
	}

	// mapの削除機能は第一引数には削除対象のmap第二引数には削除する要素のキーを指定する
	delete(mo, "name")
	fmt.Println(mo)

	// fmt.Println(m)

	// 重複排除
	duplicatedArray := []int{1, 2, 3, 3, 4, 5}
	notDuplicatedArray := []int{}

	// 空の構造体を用いることで
	myMap := map[int]struct{}{}

	for _, v := range duplicatedArray {
		_, loopValue := myMap[v]
		fmt.Println(v)
		fmt.Println(loopValue)

		if !loopValue {
			myMap[v] = struct{}{}
			notDuplicatedArray = append(notDuplicatedArray, v)
		} else {
			fmt.Println("一致したキーがあるよ")
		}
	}

	fmt.Println(notDuplicatedArray)
}
