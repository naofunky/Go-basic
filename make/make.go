package main

import "fmt"

func main() {
	// s := make([]int, 5)

	// s = append(s, 1, 2, 3)
	// fmt.Println(s)

	s := make([]int, 0, 5)
	fmt.Println(s)

	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
