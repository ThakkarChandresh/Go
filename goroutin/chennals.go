package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var a = 10

	var b = 30

	ch := make(chan int)

	go func(a, b int) {
		fmt.Println("In go routine The Sum IS:", a+b)
		ch <- a + b
	}(a, b)

	fmt.Println("Done!!!", <-ch)
}
