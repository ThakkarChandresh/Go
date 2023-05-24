package main

import (
	"fmt"
	"time"
)

func f1(n int, ch chan int) {
	time.Sleep(time.Second * 4)
	ch <- n
}

func main1() {
	/*var channel chan int

	channel <- 10 // this line will give fatal error because channel is not initialized

	c := make(chan int)

	c1 := make(<-chan string)

	c2 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", c, c1, c2)

	go f1(10, c)

	num := <-c

	fmt.Println("Received value: ", num) */

	channel := make(chan int, 4)

	//go factorial(5, channel)

	//	fmt.Println("Factorial: ", <-channel)

	/*for i := 1; i <= 20; i++ {
		go factorial(i, channel)
	} */

	go add1(channel)

	//time.Sleep(time.Second * 60)
	channel <- 100
	channel <- 200
	channel <- 300
	channel <- 400

	//go add1(channel)

	fmt.Println("Length of channel: ", len(channel))
	fmt.Println("Capacity of channel: ", cap(channel))
}

func add1(channel chan int) {
	// here channel will be in block until it receives data ?
	fmt.Println("We are inside add1 function")
	fmt.Println(234 + <-channel)

}

func factorial(n int, f chan int) {
	fact := 1

	for i := 2; i <= n; i++ {
		fact *= i
	}
	f <- fact
}
