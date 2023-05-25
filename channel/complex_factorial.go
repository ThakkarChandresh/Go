package main

import (
	"fmt"
	"runtime"
	"sync"
)

func fact(n int, c chan int) {
	res := 1

	defer func() {
		c <- res
	}()

	for i := n; i > 1; i-- {
		res *= i
	}
	// c<-res
}

func main() {
	var wg sync.WaitGroup
	//defer wg.Wait() //Writing wait in defer will never work as it executes after return!!

	wg.Add(2)
	c := make(chan int, 100)
	defer close(c)

	go func() {

		go fact(1, c)
		fmt.Println("In First", <-c)

		wg.Done()
	}()

	go func() {

		go fact(1000000, c)
		fmt.Println("In Second", <-c)

		wg.Done()
	}()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}
