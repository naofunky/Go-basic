package main

import "fmt"

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
	v.speed -= v.humanPower * 5
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

func main() {
	v := &vehicle{0, 5}
	NewVehicle(v)

	b := &bicycle{0, 5}
	NewBicycle(b)
}
