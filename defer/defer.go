package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("ファイルが指定されていません。")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// fmt.Println("ファイルを指定できました")
	fmt.Println(f)
	log.SetFlags(log.Lshortfile)
}
