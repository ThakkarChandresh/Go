package main

import "fmt"

func main() {
	var slice []int

	fmt.Println(len(slice), cap(slice), slice)
}
