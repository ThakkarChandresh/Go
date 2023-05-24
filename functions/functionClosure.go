package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	fmt.Print(a, ", ")
	fmt.Print(b, ", ")
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), ", ")
	}
}
