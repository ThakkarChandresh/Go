package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello World")

	var a = 10

	var b = 30

	var waitGroup = sync.WaitGroup{}

	waitGroup.Add(1)

	go func(a, b int) {
		fmt.Println(a + b)
		waitGroup.Done()
	}(a, b)

	waitGroup.Wait()
	fmt.Println("Done!!!")
}
