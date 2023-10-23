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

	// s4 := make([]int, 0, 3)
	// s4 = append(s4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	// fmt.Println(s4, len(s4), cap(s4))

	// s5 := make([]int, 3, 6)
	// fmt.Println(s5, len(s5), cap(s5))
	// s6 := s5[1:3]
	// fmt.Println(s6, len(s6), cap(s6))
	// s6[1] = 100
	// fmt.Println(s5, len(s5), cap(s5))
	// fmt.Println(s6, len(s6), cap(s6))

	// s7 := []int{1, 2, 3, 4, 5}
	// s8 := s7[2:5]
	// fmt.Println(s7, len(s7), cap(s7))
	// fmt.Println(s8, len(s8), cap(s8))

	// s9 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// s10 := s9[2:6]
	// fmt.Println(s9, len(s9), cap(s9))
	// fmt.Println(s10, len(s10), cap(s10))

	s11 := []int{11, 12, 13, 14, 15, 16, 17, 18} // 8,8
	s12 := s11[2:6]                              // 4,6
	fmt.Println(s11, len(s11), cap(s11))
	fmt.Println(s12, len(s12), cap(s12))
}
