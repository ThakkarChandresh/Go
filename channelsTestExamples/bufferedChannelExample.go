package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 = sync.WaitGroup{}

func sendData(c chan int) {
	c <- 10
	c <- 20
	c <- 30
	wg1.Done()
}

func readData(c chan int) {
	time.Sleep(time.Second * 4)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	wg1.Done()
}

func main() {
	wg1.Add(1)
	c := make(chan int)
	go sendData(c)
	go readData(c)
	wg1.Wait()
}
