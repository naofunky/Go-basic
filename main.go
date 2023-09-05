package main

import (
	"fmt"
	// "os"
	// "github.com/joho/godotenv"
)

var client *string

func main() {
	test := "test"
	var pointer *string = &test

	client = &test

	defer fmt.Println(test)
	fmt.Println(pointer)
	fmt.Println(*client)
}
