package main

import (
	"fmt"
)

// コンストレンズ
type customConstraints interface {
	~int | ~int16 | float32 | float64 | string
}

type NewInt int
type NewInt16 int16

// 型パラメータを持つ関数
func add[T customConstraints](x, y T) T {
	return x + y
}
func main() {
	fmt.Printf("%v\n", add(1, 2))

	var i1, i2 NewInt = 3, 4
	var i3, i4 NewInt16 = 16, 23
	fmt.Printf("%v\n", add(i1, i2))
	fmt.Printf("%v\n", add(i3, i4))
}
