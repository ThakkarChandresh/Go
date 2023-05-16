package main

import "fmt"

func main() {

	var age int = 20

	fmt.Println(age)

	var name = "Chandresh" //Type Inference

	// fmt.Println(name)
	_ = name //using blank identifier to avoid compile time error

	s := "Thakkar"
	fmt.Println(s)

	// name:= "Hello" //can't do such
	name = "Hello"
}
