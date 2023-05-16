package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var a, b, c = 10, 20, "yash"

	const PI = 3.14

	// or

	var (
		name    = "yash"
		age     = 21
		married = false
	)

	fmt.Println(a+b, c)

	fmt.Println(fmt.Sprintf(" formatting %v", a))

	fmt.Println(name+" ", age, " ", married)
}
