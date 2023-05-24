package main

import "fmt"

func main() {
	var ch chan int

	fmt.Println(ch) // this will print nil

	ch = make(chan int)

	fmt.Println(ch) // this will print address

	c := make(chan int)

	// <- send operation
	c <- 10

	// receive operation
	num := <-c

	_ = num

	fmt.Println(num)
}
