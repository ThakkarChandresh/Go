package main

import "fmt"

func main() {
	channel1 := make(chan bool)

	fmt.Println("task begin")

	channel1 <- true

	fmt.Println("task completed")
}
