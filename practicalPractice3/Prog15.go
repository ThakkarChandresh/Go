package main

import "fmt"

func main() {
HELLO:
	fmt.Println("hello")
	goto HELLO
}
