package main

import "fmt"

func main() {
	c := make(chan int, 3)

	for i := 0; i < 5; i++ {
		fmt.Println("Before Sending the data!")
		c <- i
		fmt.Println("After Sending the Data!")
	}
}
