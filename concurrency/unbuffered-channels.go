package main

import (
	"fmt"
)

func task2(channel chan int) {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	channel <- 1

	channel <- 2
}

func task3(channel chan int) {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	channel <- 3

	channel <- 4
}

func main() {

	counter := 0
	for {

		channel := make(chan int)

		go task2(channel)

		go task3(channel)

		var arr = []int{0, 0, 0, 0}

		arr[0] = <-channel
		arr[1] = <-channel
		arr[2] = <-channel
		arr[3] = <-channel

		if arr[0] == 1 {
			fmt.Println(arr)
			break
		}
		counter++
	}
	fmt.Println("counter: ", counter)
}
