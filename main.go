package main

import "fmt"

func main() {
	pointerUser := &User{
		name: "tio",
	}
	fmt.Println(pointerUser.name)

	pointerUser.ChangeNameByAtaiReceiver("iis")
	fmt.Println(pointerUser.name)

	pointerUser.ChangeNameByPointerReceiver("iis")
	fmt.Println(pointerUser.name)
}

type User struct {
	name string
}

func (u User) ChangeNameByAtaiReceiver(name string) {
	u.name = name
}

func (u *User) ChangeNameByPointerReceiver(name string) {
	u.name = name
}
