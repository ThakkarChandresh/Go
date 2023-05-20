package main

import (
	"fmt"
	"sync"
)

var countdown = sync.WaitGroup{}

func a() {
	fmt.Println("a")
	b()
}

func b() {
	fmt.Println("b")
	countdown.Add(1)
	go a()
}

func main() {
	a()
	countdown.Wait()
}
