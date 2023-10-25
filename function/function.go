package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

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
}
