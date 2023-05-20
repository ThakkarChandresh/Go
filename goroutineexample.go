package main

import (
	"fmt"
	"sync"
)

var wait = sync.WaitGroup{}

func hello() {
	defer wait.Done()
	fmt.Println("sayinghello function is executed")
}

func main() {

	wait.Add(1)
	go hello()
	wait.Wait()

}
