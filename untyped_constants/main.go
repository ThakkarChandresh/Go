package main

import "fmt"

func main() {
	const a float64 = 5.1 //typed constants
	const b = 4           //un-typed constants

	const mul float64 = a * b

	const str = "Hello " + "Go!"
	fmt.Println(str)

	const d = 5 > 10
	fmt.Println(d)

	const z int = 5
	const y float64 = 2.2 * z //this is strongly typed programming language

	const x = 5
	const y1 = 2.2 * x

	var i int = x
	var j float64 = x
	var k byte = x

	fmt.Println(i, j, k)

	const q int = 10
	const w float64 = q

	var n = 10
	var m float64 = n

}
