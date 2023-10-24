package main

import "fmt"

// 構造体の定義
type User struct {
	Name string `json:"name"`
}

func main() {

	user1 := User{
		Name: "般若",
	}
	fmt.Println("①", user1.Name)
	changeName(user1)
	fmt.Println("②", user1.Name)
	changePointerName(&user1)
	fmt.Println("③", user1.Name)

	user2 := &User{
		Name: "舐達磨",
	}
	fmt.Println("④", user2.Name)
	changeName(user1)
	fmt.Println("⑤", user2.Name)
	changePointerName(user2)
	fmt.Println("⑥", user2.Name)
}

func changeName(u User) {
	u.Name = "りん"
}

func changePointerName(u *User) {
	u.Name = "りん"
}
