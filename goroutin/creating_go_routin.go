package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var a = 10

	var b = 30
	
	go func(a, b int) {
		fmt.Println(a + b)
	}(a, b)

	fmt.Println("Done!!!")
}
