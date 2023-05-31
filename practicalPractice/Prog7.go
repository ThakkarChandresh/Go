package main

import "fmt"

func main() {
	x := 0
	arr := []int{2, 3, 5}
	x, arr = arr[0], arr[1:]
	fmt.Println(x, arr)
}
