package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup

	n := 0

	wg.Add(gr * 2)

	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			mutex.Lock()
			n++
			mutex.Unlock()
			wg.Done()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			mutex.Lock()
			n--
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(n)
}
