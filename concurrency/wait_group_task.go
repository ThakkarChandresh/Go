package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("GR1 Completed!")
		wg.Done()
	}()

	go func() {
		wg.Done()
	}()

	wg.Wait()
}
