package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("Not Blocked Baby!")
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int, 3)
	quit := make(chan int)
	go func() {
		time.Sleep(time.Second * 10)
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			time.Sleep(time.Second * 3)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
