package main

import "fmt"

func main() {
	var x = 4
	var y = 5.2

	x = int(y)

	var z = 2.
	fmt.Printf("%v ", z)
	// fmt.Println(z)

	_ = x

	// var y int
	// y = "5"

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	var x1 float64
	fmt.Println(x1 == 0.0)

	fmt.Println("The Sum is:", x+x)
}
