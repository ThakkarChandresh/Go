package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func task1() {

	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

func summation(array []int, waitgrp *sync.WaitGroup) int {
	defer (*waitgrp).Done()
	sum := 0
	for _, val := range array {
		sum += val
	}

	return sum
}

func main() {

	//runtime.GOMAXPROCS(4) // This will instruct the application to scale up on multiple cores. Here we have specified that the application can use four cores for the execution.

	wg.Add(1)

	go task1()

	var arr = []int{1, 2, 3, 4, 5, 6}

	var sum = 0
	wg.Add(2)

	go func() {
		sum += summation(arr[:len(arr)/2], &wg)
	}()
	go func() {
		sum += summation(arr[len(arr)/2:], &wg)
	}()

	wg.Wait()

	//time.Sleep(90 * time.Second)

	fmt.Println("\n", sum)

}
