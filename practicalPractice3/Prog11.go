package main

import "fmt"

func main() {
	arr := []int{3, 2, 5}
	x := 0
	x, arr = arr[0], arr[1:]
	fmt.Println(x, arr)
}
