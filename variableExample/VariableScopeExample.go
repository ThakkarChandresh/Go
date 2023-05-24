package main

import (
	"fmt"
)

var s = "Japan"

func main() {
	fmt.Println(s) // Accessing the global variable 's' and printing its value
	x := true      // Declaring and initializing a local variable 'x'

	if x { // If 'x' is true, execute the block of code inside the if statement
		y := 1 // Declaring and initializing a local variable 'y' inside the if block

		if x != false { // If 'x' is not false, execute the block of code inside the if statement
			fmt.Println(s) // Accessing the global variable 's' and printing its value
			fmt.Println(x) // Printing the value of local variable 'x'
			fmt.Println(y) // Printing the value of local variable 'y'
		}
	}
	//fmt.Println(y)

	fmt.Println(x) // Printing the value of local variable 'x'
}
