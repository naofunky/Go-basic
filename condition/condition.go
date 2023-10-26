package main

import (
	"fmt"
	// "time"
)

func main() {
	a := -1
	if a == 0 {
		fmt.Println("zero")
	} else if a > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for文は条件を省略すると無限ループになる。
	// var i int
	// for {
	// 	if i > 3 {
	// 		break
	// 	}
	// }
	// fmt.Println(i)
	// i += 1
	// time.Sleep(300 * time.Millisecond)

	// loop:
	//
	//	for i := 0; i < 10; i++ {
	//		switch i {
	//		case 2:
	//			continue
	//		case 3:
	//			continue
	//		case 8:
	//			break loop
	//		default:
	//			fmt.Printf("%v", i)
	//		}
	//		fmt.Printf("\n")
	//	}
}
