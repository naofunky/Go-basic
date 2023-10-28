package main

import "fmt"

type customConstraints interface {
	int | int16 | float32 | float64 | string
}

func add[T customConstraints](x, y T) T {
	return x + y
}
func main() {
	fmt.Printf("%v\n", add(1, 2))
}
