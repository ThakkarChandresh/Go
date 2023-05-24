package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("main execution started")
	fmt.Println("No. of CPU: ", runtime.NumCPU())
	fmt.Println("No. of goroutine: ", runtime.NumGoroutine())

	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)
	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	go f1(&waitGroup)

	fmt.Println("call to f1() method made...")

	waitGroup.Wait()

	fmt.Println("End of main")
}

func f1(waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	waitGroup.Done()
}
