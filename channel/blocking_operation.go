package main

import (
	"fmt"
	"sync"
)

func blockingFunction(channel chan int, wg *sync.WaitGroup) {
	fmt.Println(<-channel)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	ch := make(chan int)

	wg.Add(1)
	go blockingFunction(ch, &wg)

	wg.Wait() //go will smartly detect the deadlock here!

	ch <- 100
}
