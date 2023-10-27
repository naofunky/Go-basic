package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err01 := errors.New("something wrong")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)

	file := "dummy.txt"
	err3 := fileChecker(file)

	if err3 != nil {
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Printf("%v file not found\n", file)
		} else {
			fmt.Println("unknown error")
		}
	}
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in checker: %w", err)
	}
	defer f.Close()
	return nil
}
