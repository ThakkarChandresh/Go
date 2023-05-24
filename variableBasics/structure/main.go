package main

import (
	"fmt"
)

type names []string

func (n names) print() {
	for index, name := range n {
		fmt.Println("Index: ", index, " - ", name)
	}
}

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	myBook := book{"Book title", "Nikunj patel", 2023}

	fmt.Println("myBook: ", myBook)

	myBook1 := book{title: "book1 title"}

	aBook := book{title: "book1 title", author: "", year: 0}

	fmt.Println(myBook1 == aBook)

	friends := names{"Nikunj", "Rahil", "Prahlad"}

	friends.print()

	println()

	friends1 := names{"Prahlad", "Rahil", "Nikunj"}

	names.print(friends1)
}
