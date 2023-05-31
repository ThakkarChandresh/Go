package main

import (
	"fmt"
	"sync"
)

func main() {

	group := sync.WaitGroup{}

	group.Add(1)

	go func(group *sync.WaitGroup) {

		group.Done()
		group.Wait()
		group.Done()

	}(&group)

	group.Wait()
	fmt.Println("Infinite wait condition")

}
