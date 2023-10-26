package main

import (
	"fmt"
	"unsafe"
)

// メソッド一覧を登録
type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bicycle struct {
	speed      int
	humanPower int
}

// 車輌の速度を上げるメソッド
func (v *vehicle) speedUp() int {
	v.speed += v.enginePower * 10
	return v.speed
}

// 車輌の速度を下げるメソッド
func (v *vehicle) speedDown() int {
	v.speed -= v.enginePower * 5
	return v.speed
}

// 車輌の速度を上げるメソッド
func (v *bicycle) speedUp() int {
	v.speed += v.humanPower * 3
	return v.speed
}

// 車輌の速度を下げるメソッド
func (v *bicycle) speedDown() int {
	v.speed -= v.humanPower * 1
	return v.speed
}

// 車輌の速度実行関数
func NewVehicle(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

// 自転車の速度実行関数
func NewBicycle(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	// 車輌の値を登録して出力
	v := &vehicle{0, 5}
	NewVehicle(v)

	// 自転車の値を登録して出力
	b := &bicycle{0, 5}
	NewBicycle(b)

	// 下記は同じような定義をしていることを意味する
	var i1 interface{}
	var i2 any
	fmt.Printf("%[1]v %[1]T %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]v %[1]T %v\n", i2, unsafe.Sizeof(i2))

	i2 = 2

	checkType(i1)
	checkType(i2)
}
