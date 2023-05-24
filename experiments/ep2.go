package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var data = make(map[int]int)

	wg.Add(2)

	// it's reading
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println(data[i])
		}
	}()

	// it's writing
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			data[i]++
		}
	}()

	wg.Wait()
}
