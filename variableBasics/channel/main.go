package main

import (
	"fmt"
	"time"
)

/*
func main() {

	//var channel = make(chan int)

	go add()
	//channel <- 100
	//channel <- 400
	//close(channel)

	//channel <- 49
	fmt.Println("Hello")
}

func add() {
	// here channel will be in block until it receives data ?
	time.Sleep(time.Second * 5)
	//fmt.Println(234 + <-channel)
}*/

func main() {

	var channel = make(chan int)

	//go add(channel)
	//go add(channel)

	/*go func() {
		time.Sleep(time.Second * 45)
	}()*/
	//channel <- 100
	//channel <- 400

	go fib(1, channel)
	//fib(10, channel)

	for x := range channel {
		fmt.Print(x, " ")
	}

	//num := channel

	fmt.Println("Fibonacci:", channel)
	//channel <- 49
}

func fib(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
	fmt.Println("end of fib")
}

func add(channel chan int) {
	// here channel will be in block until it receives data ?
	time.Sleep(time.Second * 5)
	//fmt.Println(234 + <-channel)
}
