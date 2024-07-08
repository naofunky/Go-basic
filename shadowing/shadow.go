package main

import "fmt"

func main() {

	x := 10
	fmt.Println("outter block", x)

	if x > 5 {
		fmt.Println("inner block", x)
		a, x := 20, 5
		fmt.Println("inner block", x, a)
	}
	fmt.Println("outter block", x)
}
