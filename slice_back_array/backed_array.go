package main

import "fmt"

func main() {

	slice := make([]int, 3)

	fmt.Println(len(slice), cap(slice))

	slice = append(slice, 2)

	fmt.Println(slice)

	fmt.Println(len(slice), cap(slice))
}
