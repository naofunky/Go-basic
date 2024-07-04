package main

import "fmt"

func main() {

	x := 10
	fmt.Println("outter block", x)

	if x > 5 {
		fmt.Println("inner block", x)
		x := 5
		fmt.Println("inner block", x)
	}
	fmt.Println("outter block", x)
}
