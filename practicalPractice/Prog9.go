package main

import "fmt"

func main() {
	defer func() { fmt.Println("func-1") }()
	defer func() { fmt.Println("func-2") }()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("func-6")
		}
	}()
	defer func() {
		fmt.Println("func-3")
		panic("func-4")
	}()
	defer func() { fmt.Println("func-5") }()
}
