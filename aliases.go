package main

import "fmt"

func main() {
	var a uint8 = 127

	var b byte = 100

	// even though they have different names, byte and uit8 are the same type because they are aliases

	temp := a
	a = b
	b = temp

	fmt.Println(a, " ", b)

	type integer_type = int

	// declaring a new alias named integer_type for int
	// type alias_name = type_name

	var age integer_type = 21

	fmt.Printf("Type %T\n", age)
	fmt.Printf("Value %v", age)

}
