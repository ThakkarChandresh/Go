package main

import "fmt"

func sub(x, y int) int {
	return x - y
}

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Add 10+30: ", add(10, 30))
	fmt.Println("Subtract 10-30: ", sub(10, 30))
}
