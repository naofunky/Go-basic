package main

import (
	"fmt"
	"log"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Bob", "lob", 27},
		{"Tom", "aoi", 13},
		{"Roberto", "kei", 54},
	}

	fmt.Println("sort前", people)

	// peopleのデータがなかったらエラーを出す
	if len(people) == 0 {
		log.Fatal("データがありません。")
	}

	// ソートする
	sort.Slice(people, func(i int, j int) bool {
		return people[i].FirstName < people[j].FirstName
	})

	fmt.Println("sort後", people)
}
