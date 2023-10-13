package slice

import "fmt"

func Slice() {
	var a [3]int
	var a1 = [3]int{1, 2, 3}
	a2 := [...]int{1, 2, 3, 4, 5}

	var s1 []int
	s2 := []int{}

	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(len(a2), cap(a2))
	fmt.Println(s1)
	fmt.Println(s2)
}
