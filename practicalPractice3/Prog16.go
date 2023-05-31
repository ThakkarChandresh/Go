package main

import "fmt"

func main() {
	c := make(chan bool)
	fmt.Println("1")
	c <- true
	fmt.Println("2")
}
