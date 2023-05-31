package main

import (
	"fmt"
)

func main() {

	defer func() { fmt.Println("function 1") }()

	defer func() { fmt.Println("function 2") }()

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("function 6")
		}

	}()

	defer func() {
		fmt.Println("function 3")
		panic("function 4")
	}()

	defer func() { fmt.Println("function 5") }()

}
