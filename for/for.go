package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 { // 		fmt.Println(i, "3でも5でも割り切れる数字")
			continue
		}

		if 1%3 == 0 {
			fmt.Println(i, "3で割り切れる数字")
			continue
		}

		if i%5 == 0 {
			fmt.Println(i, "5で割り切れる数字")
			continue
		}

		// 	fmt.Println(i)
		// }

		// for i := 1; i <= 100; i++ {
		// 	if i%3 == 0 {
		// 		fmt.Println(i, "3で割り切れる数")
		// 		if i%5 == 0 {
		// 			fmt.Println(i, "3でも5でも割り切れる数")
		// 		}
		// 	} else if i%5 == 0 {
		// 		fmt.Println(i, "5で割り切れる数")
		// 	} else {
		// 		fmt.Println(i)
		// 	}
		// }

		x := []string{"北花田", "河内松原", "泉佐野", "泉南"}

		for i, v := range x {
			fmt.Println(i, v)
		}
	}
}
