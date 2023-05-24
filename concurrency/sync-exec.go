package main

import (
	"fmt"
	"sync"
)

func main() {

	var weightgroup sync.WaitGroup

	weightgroup.Add(2)

	var counter = 0

	var data = map[int]string{0: "Yash", 1: "Mihir"}

	var channel = make(chan interface{})

	go func() {

		defer weightgroup.Done()

		var receivedData = <-channel

		dp := receivedData.(map[int]string)

		dp[0] = "Dhaval"

		fmt.Println("------------", receivedData)

		for i := 0; i < 10000; i++ {
			counter--
		}
	}()

	go func() {

		defer weightgroup.Done()

		for i := 0; i < 10000; i++ {
			counter++
		}

		channel <- data
	}()

	weightgroup.Wait()

	fmt.Println(counter)

}
