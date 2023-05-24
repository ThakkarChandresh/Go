package main

import "fmt"

func goIncrement(c chan int, x int) {
	c <- x + 1
}

func main() {
	c := make(chan int)
	go goIncrement(c, 10)
	go goIncrement(c, 30)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
