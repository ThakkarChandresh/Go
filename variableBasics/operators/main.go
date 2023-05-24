package main

import "fmt"

func main() {

	var x, y = 35, 7

	fmt.Println("Addition: ", x+y)
	fmt.Println("Subtraction: ", x-y)
	fmt.Println("Multiplication: ", x*y)
	fmt.Println("Division: ", x/y)
	fmt.Println("modulo: ", x%y)

	x++
	fmt.Println("post x: ", x)

}
