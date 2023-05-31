package main

import "fmt"

func main() {
	channel := make(chan bool, 2)
	fmt.Println("task begin")
	channel <- true
	channel <- true
	fmt.Println("task completed")
}
