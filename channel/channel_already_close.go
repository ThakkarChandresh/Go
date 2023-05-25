package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func buftask1(channel chan int) {

	time.Sleep(5 * time.Second)

	channel <- 10
	channel <- 20
	channel <- 30

	close(channel)

}

func buftask2(channel chan int) {
	defer wg.Done()
	time.Sleep(10 * time.Second)

	channel <- 40
	channel <- 50
	channel <- 60

}

func main() {
	wg.Add(1)
	channel := make(chan int, 5)

	go buftask1(channel)

	go buftask2(channel)

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	wg.Wait()
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}
