package main

import "fmt"

func main() {
	ch := make(chan bool, 2)
	fmt.Println("1")
	ch <- true
	fmt.Println("2")
}
