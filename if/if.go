package main

import "fmt"

func main() {
	x := 10
	fmt.Println("outer block before if:", x) // 10

	if y := x + 1; y > 10 {
		fmt.Println("if block:", x, y) // 10 11
	} else {
		fmt.Println("else block:", x, y) // 10 11
		z := y - 1
		fmt.Println("else block:", x, y, z) // 10 11 10
	}

	fmt.Println("outer block after if:", x) // 10
	// fmt.Println(y) // Error: undefined: y
	// fmt.Println(z) // Error: undefined: z
}
