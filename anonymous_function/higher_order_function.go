package main

import "fmt"

type First func(int) int
type Second func(int) First

func sum(x, y int) int {
	return x + y
}
func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func cubeSum(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x*x + y*y*y + z*z*z
		}
	}
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}
func main() {
	partial := partialSum(3)
	fmt.Println(partial(7))

	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum(5)(6)(7))
}
