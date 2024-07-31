package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func funcDefer() {
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("Hello world")
}

// ファイルの拡張子を.htmlに変更する関数
func removeExtension(files ...string) []string {
	// 初期化
	fileNameOut := make([]string, 0, len(files))

	// 一つずつみていき、ファイル名を取り除く
	for _, v := range files {
		fileName := strings.TrimSuffix(v, path.Ext(v))
		fileNameOut = append(fileNameOut, fileName+".html")
	}
	return fileNameOut
}

// ファイルが存在するかどうか検証する関数
func fileChecker(name string) (string, error) { //複数の返り値を設定したい場合はカッコで括り指定する
	f, err := os.Open(name)
	if err != nil {
		return "", errors.New("file no found")
	}
	defer f.Close()
	return name, nil
}

// ファイル名を受け取り指定の拡張子を与える関数
func addExc(f func(file string) string, name string) {
	fmt.Println(f(name))
}

// 無名関数を引数にとる関数
func sas() func(int) int {
	return func(n int) int {
		return n * 10000
	}
}

// クロージャ
func countUp() func(int) int {
	count := 0
	return func(n int) int {
		count += n
		return count
	}
}

// ブランクreturnの例
func add(a int, b int) (sum int) {
	sum = a + b
	return sum
}

func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("0での除算はできない")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	funcDefer()
	// スライスの値を引数に渡すことで、拡張子を除いてファイル名だけを返してくれる
	files := []string{"file1.csv", "file2.csv", "file3.png"}
	fmt.Println(removeExtension(files...))

	name, err := fileChecker("main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	// 即座に実行される無名関数
	// i := 1
	// func(i int) {
	// 	fmt.Println(i)
	// }(i)

	// 任意のタイミングで関数を実行する時は、変数に代入したりする
	// f1 := func(i int) int {
	// 	return i + 1
	// }
	// fmt.Println(f1(i))

	// 無名関数を関数の引数として渡す
	f2 := func(file string) string {
		return file + ".csv"
	}
	addExc(f2, "file1")

	f3 := sas()
	fmt.Println(f3(2))

	people := []Person{
		{"夏生", 20},
		{"楠葉", 23},
		{"月香", 18},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(2)
		fmt.Printf("%v\n", v)
	}

	result, numerator, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, numerator)

	fmt.Println(add(2, 4))
}
