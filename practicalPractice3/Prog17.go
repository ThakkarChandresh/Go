package main

import "fmt"

func main() {
	arr := [2]int{2, 3}
	arr1 := [...]int{2, 3}
	fmt.Println(arr == arr1)
}
