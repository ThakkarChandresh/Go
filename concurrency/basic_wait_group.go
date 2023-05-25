package main

import (
	"fmt"
	"sync"
	"time"
)

func print100(wg *sync.WaitGroup) {
	for i := 1; i <= 100; i++ {
		fmt.Println("Second Independent Task", i)
		time.Sleep(1 * time.Microsecond)
	}
	(*wg).Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go print100(&wg)

	for i := 1; i <= 100; i++ {
		fmt.Println("Main Goroutine Independent Task", i)
	}

	wg.Wait()
}
