package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// コンストレンズ
type customConstraints interface {
	~int | ~int16 | ~float32 | ~float64 | ~string
}

type NewInt int
type NewInt16 int16
type NewFloat32 float32
type NewFloat64 float64
type NewString string

// 型パラメータを持つ関数
func add[T customConstraints](x, y T) T {
	return x + y
}

// 値を小さい方を返す関数を作成
// func min[T constraints.Ordered](x, y T) T {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// mapのgenericsで合計値を出す関数
func sumValues[K uint | string, T constraints.Float | constraints.Integer | string](m map[K]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	// fmt.Printf("%v\n", add(1, 2))
	// fmt.Printf("%v\n", add(1.2, 2.3))
	// fmt.Printf("%v\n", add(3.343, 6.789))
	// fmt.Printf("%v\n", add("Hello", "Golang"))

	// var i1, i2 NewInt = 3, 4
	// var i3, i4 NewInt16 = 16, 23
	// var i5, i6 NewFloat32 = 2.32, 4.56
	// var i7, i8 NewFloat64 = 2.32567686, 4.56768970467
	// var i9, i10 NewString = "Hello", "World"

	// fmt.Printf("%v\n", add(i1, i2))
	// fmt.Printf("%v\n", min(i1, i2))
	// fmt.Printf("%v\n", add(i3, i4))
	// fmt.Printf("%v\n", add(i5, i6))
	// fmt.Printf("%v\n", add(i7, i8))
	// fmt.Printf("%v\n", add(i9, i10))

	// mapで扱った値をgenericsを用いて、合計する
	map1 := map[string]uint{
		"B": 10,
		"A": 4,
		"C": 8,
	}

	map2 := map[uint]float32{
		1: 2.2,
		5: 6.7,
		9: 1.1,
	}

	map3 := map[uint]string{
		1: "H",
		2: "e",
		3: "l",
		4: "l",
		5: "o",
	}

	fmt.Printf("%v\n", sumValues(map1))
	fmt.Printf("%v\n", sumValues(map2))
	fmt.Printf("%v\n", sumValues(map3))
}
