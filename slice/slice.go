package main

import "fmt"

func main() {
	sl := []int{}

	// sl = append(sl, 1)
	// sl = append(sl, 2)
	sl = append(sl, 1, 2, 3, 4, 5)

	for _, v := range sl {
		fmt.Println(v)
	}

	fmt.Println(sl)
	fmt.Println(sl[0])
	fmt.Println(len(sl))
}
