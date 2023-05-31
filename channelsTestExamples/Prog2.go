package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)

	go function(ch)

	//ch <- 897876

	fmt.Println("main", <-ch)

	time.Sleep(6 * time.Second)
	//channel <- 49
}

func function(ch chan int) {
	fmt.Println("inside function")
	ch <- 79865

	fmt.Println("chandu")
	fmt.Println("<- ch: gr1: ", <-ch)
	fmt.Println("out")
	fmt.Println(<-ch)
}
