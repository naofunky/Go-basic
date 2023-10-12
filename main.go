package main

import "fmt"

// "os"
// "github.com/joho/godotenv"

var client *string

func main() {
	// test := "test"
	// var pointer *string = &test

	// client = &test

	// defer fmt.Println(test)
	// fmt.Println(pointer)
	// fmt.Println(*client)
	forname := User{name: "test"}
	getName(forname)
	fmt.Println(forname.name)

	// forname := &User{name: "test"}
	// getName(forname)
	// fmt.Println(forname.name)

	s := []string{"a", "b", "c"}
	changeSlice(s)
	fmt.Println(s)
}

type User struct {
	name string
}

func getName(u User) {
	u.name = "test2"
}

func changeSlice(s []string) {
	s[0] = "d"
}
