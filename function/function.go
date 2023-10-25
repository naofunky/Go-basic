package main

import (
	"fmt"
	"path"
	"strings"
)

func funcDefer() {
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("Hello world")
}

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

func main() {
	funcDefer()
	// スライスの値を引数に渡すことで、拡張子を除いてファイル名だけを返してくれる
	files := []string{"file1.csv", "file2.csv", "file3.png"}
	fmt.Println(removeExtension(files...))
}
