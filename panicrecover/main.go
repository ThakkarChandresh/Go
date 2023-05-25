package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("First")
	}()

	defer func() {
		fmt.Println("Second")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd! Third")
		}
	}()

	defer func() {
		fmt.Println("Fourth")
	}()

	defer func() {
		fmt.Println("Fifth")
	}()

	panic("Some Error Occurred!")
}
