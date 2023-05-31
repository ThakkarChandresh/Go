package main

import "fmt"

func main() {
	arr1 := [2]int{2, 4}
	arr2 := [...]int{2, 4}
	fmt.Println(arr1 == arr2)
}
