package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func buftask1(channel chan int) {

	channel <- 10
	channel <- 20
	channel <- 30

	wg1.Done()
}

func buftask2(channel chan int) {

	time.Sleep(time.Second * 2)

	channel <- 40
	channel <- 50
	//channel <- 60

	fmt.Println(len(channel))
	wg1.Done()
}

func main() {

	wg1.Add(2)
	channel := make(chan int, 5)

	go buftask2(channel)
	go buftask1(channel)

	//fmt.Println(<-channel)
	//fmt.Println(<-channel)
	//fmt.Println(<-channel)
	//fmt.Println(<-channel)
	//fmt.Println(<-channel)
	wg1.Wait()
}
