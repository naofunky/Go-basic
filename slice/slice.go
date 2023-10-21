package main

import "fmt"

func main() {
	// sl := []int{}

	// sl = append(sl, 1)
	// sl = append(sl, 2)
	// sl = append(sl, 1, 2, 3, 4, 5)

	// for _, v := range sl {
	// 	fmt.Println(v)
	// }

	// var nilSlice []int
	// emptySlice := []int{}

	// fmt.Println(nilSlice == nil)
	// fmt.Println(emptySlice == nil)

	// fmt.Println(len(nilSlice))
	// fmt.Println(len(emptySlice))

	// fmt.Println(cap(nilSlice))
	// fmt.Println(cap(emptySlice))

	// fmt.Println(sl)
	// fmt.Println(sl[0])
	// fmt.Println(len(sl))

	s4 := make([]int, 0, 3)
	s4 = append(s4, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(s4, len(s4), cap(s4))
}
