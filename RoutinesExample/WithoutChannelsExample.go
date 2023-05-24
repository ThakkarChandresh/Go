package main

import "fmt"

func increment(x int) int {
	return x + 1
}

func main() {
	a := increment(10)
	go increment(a)
	fmt.Println(a)
}
