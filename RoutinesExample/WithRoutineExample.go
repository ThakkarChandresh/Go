package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}

func gosay(something string) {
	for i := 0; i < 5; i = i + 1 {
		time.Sleep((3000) * time.Microsecond)
		fmt.Println(something, i)
	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(2)
	go gosay("Hello")
	gosay("Bye")
	waitGroup.Wait()
}
