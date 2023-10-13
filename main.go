package main

import (
	"encoding/json"
	"fmt"
	"go-basices/pointer"
)

// "go-basices/slice"

func main() {
	// slice.Slice()
	user := pointer.NewUser()
	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonUser))
}
